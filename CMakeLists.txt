# cmake -G "MinGW Makefiles" . && make

if(CMAKE_C_COMPILER_ID MATCHES "GNU|Clang")
  set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -fpermissive")
  set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -fpermissive")
endif()

cmake_minimum_required(VERSION 3.10)

add_definitions()

set(BLAKE_SRC
    blake2b.cpp
    sph_blake.cpp
    blake_hash.cpp
)

add_library(blake STATIC ${BLAKE_SRC})
target_link_libraries(blake)
