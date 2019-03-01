package main

// http://man7.org/linux/man-pages/man7/signal.7.html

/**
Where three values are given, the first one is usually
valid for alpha and sparc, the middle one for x86, arm, and most
other architectures, and the last one for mips.
*/

// IGHUP        1       Term    Hangup detected on controlling terminal
//                                      or death of controlling process
// SIGINT        2       Term    Interrupt from keyboard
// SIGQUIT       3       Core    Quit from keyboard
// SIGILL        4       Core    Illegal Instruction
// SIGABRT       6       Core    Abort signal from abort(3)
// SIGFPE        8       Core    Floating-point exception
// SIGKILL       9       Term    Kill signal
// SIGSEGV      11       Core    Invalid memory reference
// SIGPIPE      13       Term    Broken pipe: write to pipe with no
// 			     readers; see pipe(7)
// SIGALRM      14       Term    Timer signal from alarm(2)
// SIGTERM      15       Term    Termination signal
// SIGUSR1   30,10,16    Term    User-defined signal 1
// SIGUSR2   31,12,17    Term    User-defined signal 2
// SIGCHLD   20,17,18    Ign     Child stopped or terminated
// SIGCONT   19,18,25    Cont    Continue if stopped
// SIGSTOP   17,19,23    Stop    Stop process
// SIGTSTP   18,20,24    Stop    Stop typed at terminal
// SIGTTIN   21,21,26    Stop    Terminal input for background process
// SIGTTOU   22,22,27    Stop    Terminal output for background process
// IGBUS      10,7,10     Core    Bus error (bad memory access)
// SIGPOLL                 Term    Pollable event (Sys V).
// 			       Synonym for SIGIO
// SIGPROF     27,27,29    Term    Profiling timer expired
// SIGSYS      12,31,12    Core    Bad system call (SVr4);
// 			       see also seccomp(2)
// SIGTRAP        5        Core    Trace/breakpoint trap
// SIGURG      16,23,21    Ign     Urgent condition on socket (4.2BSD)
// SIGVTALRM   26,26,28    Term    Virtual alarm clock (4.2BSD)
// SIGXCPU     24,24,30    Core    CPU time limit exceeded (4.2BSD);
// 			       see setrlimit(2)
// SIGXFSZ     25,25,31    Core    File size limit exceeded (4.2BSD);
// SIGIOT         6        Core    IOT trap. A synonym for SIGABRT
// SIGEMT       7,-,7      Term    Emulator trap
// SIGSTKFLT    -,16,-     Term    Stack fault on coprocessor (unused)
// SIGIO       23,29,22    Term    I/O now possible (4.2BSD)
// SIGCLD       -,-,18     Ign     A synonym for SIGCHLD
// SIGPWR      29,30,19    Term    Power failure (System V)
// SIGINFO      29,-,-             A synonym for SIGPWR
// SIGLOST      -,-,-      Term    File lock lost (unused)
// SIGWINCH    28,28,20    Ign     Window resize signal (4.3BSD, Sun)
// SIGUNUSED    -,31,-     Core    Synonymous with SIGSYS

var Signal map[string]string = map[string]string{
	"1":  "SIGHUP",
	"2":  "SIGINT",
	"3":  "SIGQUIT",
	"4":  "SIGILL",
	"6":  "SIGABRT",
	"8":  "SIGFPE",
	"9":  "SIGKILL",
	"11": "SIGSEGV",
	"13": "SIGPIPE",
	"14": "SIGALRM",
	"15": "SIGTERM",
	"10": "SIGUSR1",
	"12": "SIGUSR2",
	"17": "SIGCHLD",
	"18": "SIGCONT",
	"19": "SIGSTOP",
	"20": "SIGTSTP",
	"21": "SIGTTIN",
	"22": "SIGTTOU",
	"7":  "IGBUS",
	"27": "SIGPROF",
	"31": "SIGSYS",
	"5":  "SIGTRAP",
	"23": "SIGURG",
	"26": "SIGVTALRM",
	"24": "SIGXCPU",
	"25": "SIGXFSZ",
	"16": "SIGSTKFLT",
	"29": "SIGIO",
	"30": "SIGPWR",
	"28": "SIGWINCH",
}
