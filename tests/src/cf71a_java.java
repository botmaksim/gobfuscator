import java.util.Scanner;

public class cf71a_java {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        if (!sc.hasNextInt()) return;
        int n = sc.nextInt();
        for (int i = 0; i < n; i++) {
            String s = sc.next();
            if (s.length() > 10) {
                System.out.println(s.charAt(0) + "" + (s.length() - 2) + s.charAt(s.length() - 1));
            } else {
                System.out.println(s);
            }
        }
        sc.close();
    }
}
