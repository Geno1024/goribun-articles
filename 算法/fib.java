package com.mainsoft.test;
    public class Test01 {
	static long index;
	public static void main(String[] args) {
		fibProxy(10);
		fibProxy(20);
		fibProxy(30);
		fibProxy(40);
		fibProxy(50);
	}
	public static long fib(long n) {
		index++;
		if (n <= 2) {
			return 1;
		} else {
			return fib(n - 2) + fib(n - 1);
		}
	}
	public static void fibProxy(long n) {
		long beforeTime = System.currentTimeMillis();
		long num = fib(n);
		long afterTime = System.currentTimeMillis();
		long time = afterTime - beforeTime;

		System.out.println("第" + n + "个数字为：" + num + "；方法调用次数为：" + index
				+ "；运行时间为：" + time + "毫秒。");
		index = 0;
	}
    }