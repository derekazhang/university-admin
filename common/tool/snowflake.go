package tool

// twitter 雪花算法
// 把时间戳,工作机器ID, 序列号组合成一个 64位 int
// 第一位置零, [2,42]这41位存放时间戳,[43,52]这10位存放机器id,[53,64]最后12位存放序列号

import (
	"fmt"
	tnet "github.com/toolkits/net"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	machineID     int64 // 机器 id 占10位, 十进制范围是 [ 0, 1023 ]
	sn            int64 // 序列号占 12 位,十进制范围是 [ 0, 4095 ]
	lastTimeStamp int64 // 上次的时间戳(毫秒级), 1秒=1000毫秒, 1毫秒=1000微秒,1微秒=1000纳秒
	lock          sync.Mutex
)

func init() {
	lastTimeStamp = time.Now().UnixNano() / 1000000
	InitMachineId()
}

func InitMachineId() int64 {
	var mid int64
	addrs, err := tnet.IntranetIP()
	if err != nil {
		mid = time.Now().UnixNano()
	}
	if len(addrs) != 0 {
		ipaddr := strings.Split(addrs[0], ".")
		if len(ipaddr) > 2 {
			ipaddrInt64, err := strconv.ParseInt(ipaddr[3], 10, 64)
			if err == nil {
				mid = ipaddrInt64
			} else {
				mid = time.Now().UnixNano()
			}
		} else {
			mid = time.Now().UnixNano()
		}
	} else {
		mid = time.Now().UnixNano()
	}
	// 把机器 id 左移 12 位,让出 12 位空间给序列号使用
	machineID = mid << 12
	return mid
}

func GetSnowflakeId() int64 {
	lock.Lock()
	defer lock.Unlock()
	curTimeStamp := time.Now().UnixNano() / 1000000
	if lastTimeStamp == curTimeStamp {
		newsn := atomic.AddInt64(&sn, 1)
		sn = newsn
		if sn > 4095 {
			time.Sleep(time.Millisecond)
			lastTimeStamp = time.Now().UnixNano() / 1000000
			sn = 0
		}
	} else {
		sn = 0
		lastTimeStamp = curTimeStamp
	}

	// 取 64 位的二进制数 0000000000 0000000000 0000000000 0001111111111 1111111111 1111111111  1 ( 这里共 41 个 1 )和时间戳进行并操作
	// 并结果( 右数 )第 42 位必然是 0,  低 41 位也就是时间戳的低 41 位
	rightBinValue := lastTimeStamp & 0x1FFFFFFFFFF
	// 机器 id 占用10位空间,序列号占用12位空间,所以左移 22 位; 经过上面的并操作,左移后的第 1 位,必然是 0
	rightBinValue <<= 22
	return rightBinValue | machineID | sn
}

func BatchTest(num int) {
	ot := sync.Map{}
	sy := sync.WaitGroup{}
	sy.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			sn := GetSnowflakeId()
			ot.Store(sn, 1)
			sy.Done()
		}()
	}
	sy.Wait()
	m := make(map[int64]int)
	ot.Range(func(key, value interface{}) bool {
		m[key.(int64)] = value.(int)
		return true
	})
	fmt.Println(len(m))
	fmt.Println(m)
}
