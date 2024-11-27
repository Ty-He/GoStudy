cmake_minimum_required(VERSION 3.15)
project(dp C)

set(CMAKE_C_STANDARD 17)
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)
set(CMAKE_BUILD_TYPE Debug)

file(GLOB SRC ${CMAKE_CURRENT_SOURCE_DIR}/*.c)

add_executable(dp_exe ${SRC})

target_link_libraries(dp_exe PRIVATE m)
