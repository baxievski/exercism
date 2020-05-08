import java.util.Arrays;
import java.util.List;

class ResistorColor {
    int colorCode(String color) {
        List<String> validColors = Arrays.asList(colors());
        if (validColors.contains(color)) {
            return validColors.indexOf(color);
        }
        return -1;
    }

    String[] colors() {
        String[] result = {"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"};
        return result;
    }

    public static void main(String[] args) {
        System.out.println(Arrays.asList(new ResistorColor().colors()).indexOf("red"));
    }
}
