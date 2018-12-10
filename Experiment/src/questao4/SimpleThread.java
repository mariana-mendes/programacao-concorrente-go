package questao4;

import java.text.NumberFormat;

class SimpleThread extends Thread {
    public SimpleThread() {
    }
    public void run() {
        for (int i = 0; i < 10; i++) {
            try {
                sleep((int)(Math.random() * 1000));
            } catch (InterruptedException e) {}
        }
    }
}

class Main {
    public static void main(String[] args) {
    	
    	Runtime runtime = Runtime.getRuntime();

    	NumberFormat format = NumberFormat.getInstance();

    	StringBuilder sb = new StringBuilder();
    	long allocatedMemory = runtime.totalMemory();
    	sb.append("allocated memory: " + format.format(allocatedMemory / 1024) + "<br/>");
    	
    	for (int i = 0; i < 10; i++) {
        	long allocatedMemory2 = runtime.totalMemory();
        	sb.append("allocated memory: " + format.format(allocatedMemory2 / 1024) + "<br/>");
        	System.out.println(sb);
            new SimpleThread().start();
            new SimpleThread().start();
    	}
   
    	

    }
}