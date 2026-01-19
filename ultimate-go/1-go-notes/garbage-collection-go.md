As of **2026**, Golang uses a **Concurrent Mark-and-Sweep** garbage collector, which was recently upgraded with a new architecture called **"Green Tea"** (introduced in Go 1.25 as an experiment and enabled by default in Go 1.26).

Go’s GC is specifically designed for low-latency, meaning it prioritizes short "Stop-The-World" (STW) pauses over absolute maximum throughput.

---

### 1. The Core Algorithm: Tricolor Mark-and-Sweep

Historically and conceptually, the GC is built on the **Tricolor Marking** abstraction. It categorizes every object on the heap into three "colors":

* **White:** Potential garbage. These are objects that have not been reached by the scanner yet.
* **Gray:** Reachable objects that are currently being scanned (their pointers to other objects haven't been fully followed).
* **Black:** Confirmed live objects. The GC has finished scanning these and all objects they point to.

**The Process:**

1. **Mark Setup (STW):** A tiny pause to turn on "write barriers" so the GC can track changes made while it's running.
2. **Marking (Concurrent):** The GC starts from "roots" (global variables and stack pointers) and moves through the object graph, turning white objects gray, then black.
3. **Mark Termination (STW):** A second short pause to finish marking any remaining gray objects.
4. **Sweeping (Concurrent):** Reclaims the memory of all remaining white objects.

---

### 2. The 2026 Update: "Green Tea"

In **Go 1.26**, the implementation shifted from scanning individual objects to a **span-based** approach known as **Green Tea**.

* **How it differs:** Instead of jumping around the heap to scan single objects (which causes CPU cache misses), Green Tea scans entire **8 KiB memory spans** at once.
* **Benefits:** It significantly improves cache locality and reduces the CPU overhead of the GC by **10% to 40%**.
* **Modern Hardware:** It is specifically optimized to use vector instructions (like AVX on Intel/AMD) to scan memory faster than previous versions.

---

### 3. Key Characteristics of Go’s GC

Unlike languages like Java or C#, Go's garbage collector has several unique design choices:

| Feature | Go's Approach | Why? |
| --- | --- | --- |
| **Generational** | **No** | Go relies on **Escape Analysis** to put short-lived objects on the stack instead of the heap, reducing the need for a "Young Generation" collector. |
| **Compacting** | **No** | Go does not move objects around in memory. This keeps pointer operations very fast and simplifies the runtime. |
| **Concurrency** | **High** | Most of the work happens while your code is running, keeping pauses under 1 millisecond in most cases. |
| **Tuning** | **Minimal** | Developers primarily use the `GOGC` variable (percentage of heap growth) or `GOMEMLIMIT` to control behavior. |