cmake_minimum_required(VERSION 3.10)
project(Sherk C)

set(CMAKE_C_STANDARD 99)

# windows 的 Cmake 编译出错，添加此即可
link_libraries(ws2_32)

# sherkd表示编译后的可执行文件名为 sherkd
add_executable(sherkd

        ./sherkd.c

        )