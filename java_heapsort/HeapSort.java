import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;

public class HeapSort {

  public void siftDown(byte[][] data, int lo, int hi, int first) {
    int root = lo;
    while(true){
      int child = 2*root + 1;
      if(child >= hi)
        break;
      if(child+1 < hi && less(data[first+child], data[first+child+1]))
        child++;
      if(!less(data[first+root], data[first+child]))
        return;
      swap(data, first+root, first+child);
      root = child;
    }
  }

  public boolean less(byte[] a, byte[] b){
    // I assume a and b have the same length
    int n = 0;
    while(n<a.length){
      if( ((int)a[n] & 0xFF) < ((int)b[n] & 0xFF) )
        return true;
      if( ((int)a[n] & 0xFF) > ((int)b[n] & 0xFF) )
        return false;
      n++;
    }
    return false;
  }

  public void swap(byte[][] data, int a, int b) {
    byte[] temp = data[a];
    data[a] = data[b];
    data[b] = temp;
  }

  public void heapSort(byte[][] data, int a, int b){
    int first = a;
    int lo = 0;
    int hi = b - a;
    // build max-heap from index a to index b
    for(int i=(hi-1)/2; i >= 0; i--)
      siftDown(data, i, hi, first);
    // pop maximum to the end and rebuild the former elements
    for(int i=hi-1; i >= 0; i--){
      swap(data, first, first+i);
      siftDown(data, lo, i, first);
    }
  }

  public static void main(String[] args) {
    File input = new File("java_heapsort/input");
    long len = input.length();
    byte[][] data = new byte[(int)(len/100)][100];

    // generate original data
    try(FileInputStream fis = new FileInputStream(input)){
      for(int i=0; i<data.length; i++)
        for(int j=0; j<100; j++)
          data[i][j] = (byte) fis.read();
    } catch (IOException e) {
      e.printStackTrace();
    }

    // call heapSort method
    HeapSort hs = new HeapSort();
    hs.heapSort(data, 0, data.length);

    // write in the output file
    File output = new File("java_heapsort/java_output");
    try(FileOutputStream fos = new FileOutputStream(output)){
      for(byte[] buf : data)
        fos.write(buf);
    } catch (IOException e) {
      e.printStackTrace();
    }
  }

}
