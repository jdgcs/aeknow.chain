#include <cstdarg>
#include <cstdint>
#include <cstdlib>
#include <new>

extern "C" {

void ecrecover(const unsigned char *input, unsigned char *output);

} // extern "C"
