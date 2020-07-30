## Software design

You must develop a design philosophy that establishes a set of guidelines. This is more important than developing a set of rules or patterns you apply blindly. Guidelines help to formulate, drive and validate decisions. You can't begin to make the best decisions without understanding the impact of your decisions. Every decision you make, every line of code you write comes with trade-offs.

### Open Your Mind

Technology changes quickly but people's minds change slowly.
Easy to adopt new technology but hard to adopt new ways of thinking.

### Interesting Questions - What do they mean to you?

Is it a good program?
Is it an efficient program?
Is it correct?
Was it done on time?
What did it cost?

### Legacy Software

"We think awful code is written by awful devs. But in reality, it's written by reasonable devs in awful circumstances." - Sarah Mei

"There are many reasons why programs are built the way they are, although we may fail to recognize the multiplicity of reasons because we usually look at code from the outside rather than by reading it. When we do read code, we find that some of it gets written because of machine limitations, some because of language limitations, some because of programmer limitations, some because of historical accidents, and some because of specifications—both essential and inessential. - Gerald M. Weinberg"

### Mental Models

You must constantly make sure your mental model of your projects are clear. When you can't remember where a piece of logic is or you can't remember how something works, you are losing your mental model of the code. This is a clear indication that refactoring is a must. Focus time on structuring code that provides the best mental model possible and code review for this as well.

### Correctness vs Performance

You want to write code that is optimized for correctness. Don't make coding decisions based on what you think might perform better. You must benchmark or profile to know if code is not fast enough. Then and only then should you optimize for performance. This can't be done until you have something working.

Improvement comes from writing code and thinking about the code you write. Then refactoring the code to make it better. This requires the help of other people to also read the code you are writing. Prototype ideas first to validate them. Try different approaches or ask others to attempt a solution. Then compare what you have learned.

Too many developers are not prototyping their ideas first before writing production code. It is through prototyping that you can validate your thoughts, ideas and designs. This is the time when you can break down walls and figure out how things work. Prototype in the concrete and consider contracts after you have a working prototype.

Refactoring must become part of the development cycle. Refactoring is the process of improving the code from the things that you learn on a daily basis. Without time to refactor, code will become impossible to manage and maintain over time. This creates the legacy issues we are seeing today.

### Rules

Rules have costs.
Rules must pull their weight - Don’t be clever (high level).
Value the standard, don’t idolize it.
Be consistent!
Semantics convey ownership.

### Code Reviews

You can't look at a piece of code, function or algorithm and determine if it smells good or bad without a design philosophy. These four major categories are the basis for code reviews and should be prioritized in this order: Integrity, Readability, Simplicity and then Performance. You must consciously and with great reason be able to explain the category you are choosing.

### Integrity

We need to become very serious about reliability.

There are two driving forces behind integrity:

Integrity is about every allocation, read and write of memory being accurate, consistent and efficient. The type system is critical to making sure we have this micro level of integrity.
Integrity is about every data transformation being accurate, consistent and efficient. Writing less code and error handling is critical to making sure we have this macro level of integrity.
Write Less Code:

There have been studies that have researched the number of bugs you can expect to have in your software. The industry average is around 15 to 50 bugs per 1000 lines of code. One simple way to reduce the number of bugs, and increase the integrity of your software, is to write less code.

Bjarne Stroustrup stated that writing more code than you need results in Ugly, Large and Slow code:

Ugly: Leaves places for bugs to hide.
Large: Ensures incomplete tests.
Slow: Encourages the use of shortcuts and dirty tricks.

### Error Handling:

When error handling is treated as an exception and not part of the main code, you can expect the majority of your critical failures to be due to error handling.

48 critical failures were found in a study looking at a couple hundred bugs in Cassandra, HBase, HDFS, MapReduce, and Redis.

92% : Failures from bad error handling
35% : Incorrect handling
25% : Simply ignoring an error
8% : Catching the wrong exception
2% : Incomplete TODOs
57% System specific
23% : Easily detectable
34% : Complex bugs
8% : Failures from latent human errors

### Readability

We must structure our systems to be more comprehensible.

This is about writing simple code that is easy to read and understand without the need of mental exhaustion. Just as important, it's about not hiding the cost/impact of the code per line, function, package and the overall ecosystem it runs in.

Example Readability Issue

Code Must Never Lie

We have all been here if you have been programming long enough. At this point it doesn't matter how fast the code might be if no one can understand or maintain it moving forward.

Quotes

"This is a cardinal sin amongst programmers. If code looks like it’s doing one thing when it’s actually doing something else, someone down the road will read that code and misunderstand it, and use it or alter it in a way that causes bugs. That someone might be you, even if it was your code in the first place." - Nate Finch

Code Must Never Lie

Average Developer

You must be aware of who you are on your team. When hiring new people, you must be aware of where they fall. The code must be written for the average developer to comprehend. If you are below average, you have the responsibility to come up to speed. If you are the expert, you have the responsibility to reduce being clever.

Quotes

"Can you explain it to the median user (developer)? as opposed to will the smartest user (developer) figure it out?" - Peter Weinberger (inventor of AWK)

Real Machine

In Go, the underlying machine is the real machine rather than a single abstract machine. The model of computation is that of the computer. Here is the key, Go gives you direct access to the machine while still providing abstraction mechanisms to allow higher-level ideas to be expressed.

Quotes

"Making things easy to do is a false economy. Focus on making things easy to understand and the rest will follow." - Peter Bourgon

### Simplicity

We must understand that simplicity is hard to design and complicated to build.

This is about hiding complexity. A lot of care and design must go into simplicity because this can cause more problems then good. It can create issues with readability and it can cause issues with performance.

Complexity Sells Better

Focus on encapsulation and validate that you're not generalizing or even being too concise. You might think you are helping the programmer or the code but validate things are still easy to use, understand, debug and maintain.

Quotes

"Simplicity is a great virtue but it requires hard work to achieve it and education to appreciate it. And to make matters worse: complexity sells better." - Edsger W. Dijkstra

"Everything should be made as simple as possible, but not simpler." - Albert Einstein

"You wake up and say, I will be productive, not simple, today." - Dave Cheney

Encapsulation

Encapsulation is what we have been trying to figure out as an industry for 40 years. Go is taking a slightly new approach with the package. Bringing encapsulation up a level and providing richer support at the language level.

Quotes

Paraphrasing: "Encapsulation and the separation of concerns are drivers for designing software. This is largely based on how other industries handle complexity. There seems to be a human pattern of using encapsulation to wrestle complexity to the ground." - Brad Cox (inventor of Objective C)

"The purpose of abstraction is not to be vague, but to create a new semantic level in which one can be absolutely precise - Edsger W. Dijkstra

"Computing is all about abstractions. Those below yours are just details. Those above yours are limiting complicated crazy town." - Joe Beda

### Performance

We must compute less to get the results we need.

This is about not wasting effort and achieving execution efficiency. Writing code that is mechanically sympathetic with the runtime, operating system and hardware. Achieving performance by writing less and more efficient code but staying within the idioms and framework of the language.

Quotes

"Programmers waste enormous amounts of time thinking about, or worrying about, the speed of noncritical parts of their programs, and these attempts at efficiency actually have a strong negative impact when debugging and maintenance are considered. We should forget about small efficiencies, say about 97% of the time: premature optimization is the root of all evil. Yet we should not pass up our opportunities in that critical 3%." — Donald E. Knuth

Rules of Performance:

- Never guess about performance.
- Measurements must be relevant.
- Profile before you decide something is performance critical.
- Test to know you are correct.

Broad Engineering

Performance is important but it can't be your priority unless the code is not running fast enough. You only know this once you have a working program and you have validated it. We place those who we think know how to write performant code on a pedestal. We need to put those who write code that is optimized for correctness and performs fast enough on those pedestals.

### Micro-Optimizations

Micro-Optimizations are about squeezing every ounce of performance as possible. When code is written with this as the priority, it is very difficult to write code that is readable, simple or idiomatic. You are writing clever code that may require the unsafe package or you may need to drop into assembly.

### Design Philosophy:

Interfaces give programs structure.
Interfaces encourage design by composition.
Interfaces enable and enforce clean divisions between components.
The standardization of interfaces can set clear and consistent expectations.
Decoupling means reducing the dependencies between components and the types they use.
This leads to correctness, quality and performance.
Interfaces allow you to group concrete types by what they do.
Don't group types by a common DNA but by a common behavior.
Everyone can work together when we focus on what we do and not who we are.
Interfaces help your code decouple itself from change.
You must do your best to understand what could change and use interfaces to decouple.
Interfaces with more than one method have more than one reason to change.
Uncertainty about change is not a license to guess but a directive to STOP and learn more.
You must distinguish between code that:
defends against fraud vs protects against accidents
Validation:

Use an interface when:

users of the API need to provide an implementation detail.
API’s have multiple implementations they need to maintain internally.
parts of the API that can change have been identified and require decoupling.
Don't use an interface:

for the sake of using an interface.
to generalize an algorithm.
when users can declare their own interfaces.
if it's not clear how the interface makes the code better.

Interface And Composition Design
Design Philosophy:

Interfaces give programs structure.
Interfaces encourage design by composition.
Interfaces enable and enforce clean divisions between components.
The standardization of interfaces can set clear and consistent expectations.
Decoupling means reducing the dependencies between components and the types they use.
This leads to correctness, quality and performance.
Interfaces allow you to group concrete types by what they do.
Don't group types by a common DNA but by a common behavior.
Everyone can work together when we focus on what we do and not who we are.
Interfaces help your code decouple itself from change.
You must do your best to understand what could change and use interfaces to decouple.
Interfaces with more than one method have more than one reason to change.
Uncertainty about change is not a license to guess but a directive to STOP and learn more.
You must distinguish between code that:
defends against fraud vs protects against accidents
Validation:

Use an interface when:

users of the API need to provide an implementation detail.
API’s have multiple implementations they need to maintain internally.
parts of the API that can change have been identified and require decoupling.
Don't use an interface:

for the sake of using an interface.
to generalize an algorithm.
when users can declare their own interfaces.
if it's not clear how the interface makes the code better.
Resources:

Methods, interfaces and Embedding - William Kennedy
Composition with Go - William Kennedy
Reducing type hierarchies - William Kennedy
Interface pollution in Go - Burcu Dogan
Application Focused API Design - William Kennedy
Avoid interface pollution - William Kennedy
Interface Values Are Valueless - William Kennedy
Interface Semantics - William Kennedy

Package-Oriented Design
Package Oriented Design allows a developer to identify where a package belongs inside a Go project and the design guidelines the package must respect. It defines what a Go project is and how a Go project is structured. Finally, it improves communication between team members and promotes clean package design and project architecture that is discussable.

Learn More

Concurrent Software Design
Concurrency means “out of order” execution. Taking a set of instructions that would otherwise be executed in sequence and finding a way to execute them out of order and still produce the same result. For the problem in front of you, it has to be obvious that out of order execution would add value. When I say value, I mean add enough of a performance gain for the complexity cost. Depending on your problem, out of order execution may not be possible or even make sense.

It’s also important to understand that concurrency is not the same as parallelism. Parallelism means executing two or more instructions at the same time. This is a different concept from concurrency. Parallelism is only possible when you have at least 2 operating system (OS) and hardware threads available to you and you have at least 2 Goroutines, each executing instructions independently on each OS/hardware thread.

Both you and the runtime have a responsibility in managing the concurrency of the application. You are responsible for managing these three things when writing concurrent software:

Design Philosophy:

The application must startup and shutdown with integrity.
Know how and when every goroutine you create terminates.
All goroutines you create should terminate before main returns.
Applications should be capable of shutting down on demand, even under load, in a controlled way.
You want to stop accepting new requests and finish the requests you have (load shedding).
Identify and monitor critical points of back pressure that can exist inside your application.
Channels, mutexes and atomic functions can create back pressure when goroutines are required to wait.
A little back pressure is good, it means there is a good balance of concerns.
A lot of back pressure is bad, it means things are imbalanced.
Back pressure that is imbalanced will cause:
Failures inside the software and across the entire platform.
Your application to collapse, implode or freeze.
Measuring back pressure is a way to measure the health of the application.
Rate limit to prevent overwhelming back pressure inside your application.
Every system has a breaking point, you must know what it is for your application.
Applications should reject new requests as early as possible once they are overloaded.
Don’t take in more work than you can reasonably work on at a time.
Push back when you are at critical mass. Create your own external back pressure.
Use an external system for rate limiting when it is reasonable and practical.
Use timeouts to release the back pressure inside your application.
No request or task is allowed to take forever.
Identify how long users are willing to wait.
Higher-level calls should tell lower-level calls how long they have to run.
At the top level, the user should decide how long they are willing to wait.
Use the Context package.
Functions that users wait for should take a Context.
These functions should select on <-ctx.Done() when they would otherwise block indefinitely.
Set a timeout on a Context only when you have good reason to expect that a function's execution has a real time limit.
Allow the upstream caller to decide when the Context should be canceled.
Cancel a Context whenever the user abandons or explicitly aborts a call.
Architect applications to:
Identify problems when they are happening.
Stop the bleeding.
Return the system back to a normal state.
Index of the three part series:

Scheduling In Go : Part I - OS Scheduler
Scheduling In Go : Part II - Go Scheduler
Scheduling In Go : Part III - Concurrency
Channel Design
Channels allow goroutines to communicate with each other through the use of signaling semantics. Channels accomplish this signaling through the use of sending/receiving data or by identifying state changes on individual channels. Don't architect software with the idea of channels being a queue, focus on signaling and the semantics that simplify the orchestration required.

Language Mechanics:

Use channels to orchestrate and coordinate goroutines.
Focus on the signaling semantics and not the sharing of data.
Signaling with data or without data.
Question their use for synchronizing access to shared state.
There are cases where channels can be simpler for this but initially question.
Unbuffered channels:
Receive happens before the Send.
Benefit: 100% guarantee the signal being sent has been received.
Cost: Unknown latency on when the signal will be received.
Buffered channels:
Send happens before the Receive.
Benefit: Reduce blocking latency between signaling.
Cost: No guarantee when the signal being sent has been received.
The larger the buffer, the less guarantee.
Buffer of 1 can give you one delayed send of guarantee.
Closing channels:
Close happens before the Receive. (like Buffered)
Signaling without data.
Perfect for signaling cancellations and deadlines.
NIL channels:
Send and Receive block.
Turn off signaling
Perfect for rate limiting or short-term stoppages.
Design Philosophy:

Depending on the problem you are solving, you may require different channel semantics. Depending on the semantics you need, different architectural choices must be taken.

If any given Send on a channel CAN cause the sending goroutine to block:
Be careful with Buffered channels larger than 1.
Buffers larger than 1 must have reason/measurements.
Must know what happens when the sending goroutine blocks.
If any given Send on a channel WON'T cause the sending goroutine to block:
You have the exact number of buffers for each send.
Fan Out pattern
You have the buffer measured for max capacity.
Drop pattern
Less is more with buffers.
Don’t think about performance when thinking about buffers.
Buffers can help to reduce blocking latency between signaling.
Reducing blocking latency towards zero does not necessarily mean better throughput.
If a buffer of one is giving you good enough throughput then keep it.
Question buffers that are larger than one and measure for size.
Find the smallest buffer possible that provides good enough throughput.
