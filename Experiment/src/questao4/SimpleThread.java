package questao4;

class SimpleThread extends Thread {
	public SimpleThread() {
	}

	public void run() {

		for (int i = 0; i < 10; i++) {
			try {
				sleep((int) (Math.random() * 1000));
			} catch (InterruptedException e) {
			}
		}
	}
}

class Main {
	public static void main(String[] args) {
		Runtime runtime = Runtime.getRuntime();
		System.out.println("Diferenca: "+  (runtime.totalMemory() - runtime.freeMemory())/1024);

		System.out.println("Memoria inicial: " + runtime.totalMemory()/1024);	
		for (int j = 0; j < 12; j++) {
			for (int i = 0; i < Math.pow(2, j); i++) {
				new SimpleThread().start();
				new SimpleThread().start();
			}

			long allocatedMemory2 = runtime.totalMemory() - runtime.freeMemory();
			System.out.println("numero de threads: " + (int) Math.pow(2, j) * 2);
			System.out.println("memoria: " + allocatedMemory2 / 1024);
		}
	}
}
