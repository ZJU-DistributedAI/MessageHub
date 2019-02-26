package main

import (
	"sync"
	"sync/atomic"

	"github.com/pkg/errors"
	"github.com/ethereum/go-ethereum/rpc"
	"fmt"
	"time"
)



//TODO


type sig struct{}



type f func() ([]string,*rpc.Client)

var ans chan Receipt


type Worker struct {
	// pool who owns this worker.
	pool *Pool

	// 任务通道
	task chan f
}


type Pool struct{


	capacity int32
	running  int32
	//创建一个sig通道,用来判断是否当前没有worker空余
	freeSignal chan sig

	workers []*Worker

	release chan sig

	lock sync.Mutex

	once sync.Once

}


func NewPool(number int32)(*Pool){



	ans=make(chan Receipt)


	pool:=&Pool{
		capacity:number,
		running:0,
		freeSignal:make(chan sig),
		workers:[]*Worker{},
		release:make(chan sig),
		lock:sync.Mutex{},
		once:sync.Once{},
	}

	return pool
}




func (p *Pool) getWorker() *Worker{

	var w *Worker
	//当前运行的worker是否已经达到了运行容量上限
	waiting:=false

	//涉及到从worker队列中获取worker,因此加锁
	p.lock.Lock()

	workers:=p.workers
	n:=len(workers)-1

	//当前worker队列为空(无空闲worker)
	if n<0{
		//运行的worker已经达到数量上限，那么当前worker只能等待
		if p.running>=p.capacity{
			waiting=true
		}else{
			//否则那么该worker还可以运行(注意运行容量和worker数无关)
			p.running++
		}
	}else{
		//有空闲的worker
		<-p.freeSignal
		w=workers[n]
		workers[n]=nil
		p.workers=workers[:n]

	}
	p.lock.Unlock()

	if waiting{

		//阻塞，直到freeSignal有空闲的worker
		<-p.freeSignal
		p.lock.Lock()
		l:=len(workers)-1
		w=workers[l]
		workers[l]=nil
		p.workers=workers
		p.lock.Unlock()
	}else if w==nil{
		//当前空闲，且之前的worker都在忙，则可以新建一个worker
		w=&Worker{

			pool:p,
			task:make(chan f),
		}
		w.run()

	}


	return w
}


func (w *Worker) run(){


	go func(){

		//阻塞直到有第一个任务进入w.task通道
		for {

			f:=<-w.task
			if f==nil{
				atomic.AddInt32(&w.pool.running,-1)
				return

			}

			//本来以下的代码都要放到f()当中，但因为会产生问题，因此先保持高耦合度，待测试通过会进行封装
			txHashes,client:=f()

			//这一段可以用createGoRoutine函数替代，但还在测试当中暂且先这样
			isDone:=0
			for{
				//fmt.Println("开始不断执行")
				if isDone==len(txHashes){
					break
				}
				for i:=0;i<len(txHashes);i++{
					if txHashes[i]==""{
						continue
					}
					ans:=getTransactionReceipt(client,txHashes[i])
					if ans.BlockHash!=""{
						fmt.Println("ans txHash: ",ans.TransactionHash)
						fmt.Println("blockHash: ",ans.BlockHash)
						txHashes[i]=""
						isDone++
					}
				}

				time.Sleep(1000)
			}

			fmt.Println("结束该任务")
			w.pool.putWorker(w)


		}


	}()


}




func (w *Worker) stop(){

	w.sendTask(nil)

}

func (w *Worker) sendTask(task f){
	w.task <-task
}


//Worker回收，对goroutine进行复用
func (p *Pool) putWorker(worker *Worker){

	p.lock.Lock()
	p.workers=append(p.workers,worker)
	p.lock.Unlock()
	p.freeSignal <-sig{}
}


//提交任务到任务池

func (p *Pool) submit(task f) error{

	//协程池已经关闭
	if len(p.release)>0 {
		return errors.New("close pool fail")
	}

	w:=p.getWorker()
	w.sendTask(task)

	return nil
}


//动态控制Pool的大小
func (p *Pool) resize(size int){

	if size <p.Cap(){

		diff:=p.Cap()-size

		for i:=0;i<diff;i++{
			p.getWorker().stop()
		}
	}else if size==p.Cap(){
		return
	}

	atomic.StoreInt32(&p.capacity,int32(size))
}


func (p *Pool) Cap() (int){

	return int(p.capacity)

}




