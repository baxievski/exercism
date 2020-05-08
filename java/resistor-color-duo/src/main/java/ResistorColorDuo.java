import java.util.Arrays;
import java.util.List;

class ResistorColorDuo {
    int value(String[] colors) {
        int result = 0;
        for (int i = 0; i < 2; i++) {
            result += colorCode(colors[i]) * (int) Math.pow(10, 1-i);
        }
        return result;
    }

    int colorCode(String color) {
        List<String> validColors = Arrays.asList(colors());
        if (!validColors.contains(color)) {
            return -1;
        }
        return validColors.indexOf(color);
    }

    String[] colors() {
        String[] result = {
                "black",
                "brown",
                "red",
                "orange",
                "yellow",
                "green",
                "blue",
                "violet",
                "grey",
                "white"
        };
        return result;
    }

    public static void main(String[] args) {
        String[] c = {"blue", "grey", "red"};
        System.out.println(new ResistorColorDuo().value(c));
    }
}
