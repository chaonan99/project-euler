cmake_minimum_required(VERSION 2.8)
project(PROJECT_EULER)

set(CMAKE_C_COMPILER "g++-7")
set(CMAKE_CXX_COMPILER "g++-7")
set (CMAKE_CXX_STANDARD 11)

# Use compiler optimization or not
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -O")
add_executable(p061 junior/p061.cxx)
add_executable(p062 junior/p062.cxx)