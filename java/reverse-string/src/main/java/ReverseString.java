import java.util.List;
import java.util.Arrays;

class ReverseString {
    String reverse(String inputString) {
        if (inputString == ""){
            return inputString;
        }

        if (inputString == null) {
            return inputString;
        }

        char[] inputChars = inputString.toCharArray();
        char[] result = new char[inputChars.length];

        for (int i = 0; i <= inputChars.length - 1; i++) {
            result[i] = inputChars[inputChars.length - i - 1];
        }
        return new String(result);
    }

    public static void main(String[] args) {
        String result = new ReverseString().reverse("robot");
        System.out.println(result);
    }
}