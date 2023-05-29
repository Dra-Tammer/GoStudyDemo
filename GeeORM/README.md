# GeeORM

*在**Go**语言中，**sync.Mutex**是一种互斥锁（**Mutual Exclusion Lock**）类型，它同步多个**goroutine**的访问。
**一个互斥锁是在代码上创建一个互斥区域，该区域代码同时只能被一个**goroutine**执行。当其他**goroutine**试图访问这个互斥区域时，它们会被阻塞，直到获得锁的**goroutine**释放锁。
**在**Go**语言中，**sync.Mutex**类型的变量可以通过**Lock**和**Unlock**方法来进行加锁和解锁操作。需要注意的是，**Lock**和**Unlock**操作必须成对出现，否则会导致死锁等问题。*

[]*log.Logger{errorLog, infoLog} *//* *切片类型的变量，每个元素指向**log.logger**的指针，而**logger**又是一个结构体*

Discard是一个io.Writer接口，对它的所有Write调用都会无实际操作的成功返回。

`s.sqlVars = append(s.sqlVars, values...)`

在Go语言中，`...`被称为省略号（Ellipsis），它表示将一个slice类型的变量"打散"后追加到函数的参数列表中。在这句代码中，`values`是一个slice类型的变量，`...`表示将它展开成一个个的元素，然后追加到`s.sqlVars`中。这种语法在函数调用时也可以使用，可以方便地将一个slice传递给函数。

需要注意的是，在使用省略号作为函数参数的语法时，省略号必须放在参数列表的最后一个位置，因为它把之后的参数全部打散了。