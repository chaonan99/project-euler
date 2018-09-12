cmake_minimum_required(VERSION 2.8)
project(PROJECT_EULER)

set(CMAKE_C_COMPILER "g++-7")
set(CMAKE_CXX_COMPILER "g++-7")
set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -O")

find_package(Boost REQUIRED)
include_directories(${Boost_INCLUDE_DIRS})

# add_executable(p061 junior/p061.cxx)
add_executable(p101 junior/p101.cxx)
target_link_libraries(p101 ${Boost_LIBRARIES})
target_link_libraries(p101 gmp)
