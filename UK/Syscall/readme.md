### System call
> In computing, a system call is the programmatic way in which a computer program requests a service from the 
kernel of the operating system it is executed on.<br>
This may include hardware-related services(for example, accessing a hard disk drive), creation and executing of new  processes, and communication with integral kernal service such as process scheduling.<br>
 System calls provide an essential interface between a process and the operating system.a
 
 
#### Makeing a syscall
> syscall() saves CPU registers before making the system call, restores the registers upon return from the systemcall,
and stores any error code returned by the system call in errno(3) if an error occurs. 