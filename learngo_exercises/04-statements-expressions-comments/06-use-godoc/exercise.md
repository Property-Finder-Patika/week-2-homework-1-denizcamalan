## Solution 

- run go doc runtime NumCPU for documentation

    package runtime // import "runtime"

    func NumCPU() int
    NumCPU returns the number of logical CPUs usable by the current process.

    The set of available CPUs is checked by querying the operating system at
    process startup. Changes to operating system CPU allocation after process
    startup are not reflected.

- go doc -src runtime NumCPU for source code

    package runtime // import "runtime"

    // NumCPU returns the number of logical CPUs usable by the current process.
    //
    // The set of available CPUs is checked by querying the operating system
    // at process startup. Changes to operating system CPU allocation after
    // process startup are not reflected.

    func NumCPU() int {
            return int(ncpu)
    }