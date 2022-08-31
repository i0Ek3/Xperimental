// public means this file's name should begin with upper letter
public class Test {
    public static void main(String[] args) {
        char c = 'b';
        switch(c) {
            case 'a':
                System.out.println("This is a test!");
                break;
            case 'b':
                System.out.println("This is a test too.");
                break;
            default:
                System.out.println("unknown test.");
        }
    }
}
