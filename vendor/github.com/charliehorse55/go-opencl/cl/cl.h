#define CL_USE_DEPRECATED_OPENCL_1_2_APIS
#if defined(__APPLE__)
#   include <OpenCL/cl.h>
#   include <OpenCL/cl_ext.h>
#else
#define CL_TARGET_OPENCL_VERSION 120
#   include <CL/cl.h>
#   include <CL/cl_ext.h>
#endif
