#include <stdio.h>
#include <stdint.h>

int main(void) {
    // Hard-coded address (example only; likely invalid in real programs)
    uintptr_t addr = 0x2dc4877ac008;

    // Cast to a pointer to int
    int *p = (int *)addr;

    // Dereference (dangerous if addr is not a valid int)
    printf("%d\n", *p);

    return 0;
}