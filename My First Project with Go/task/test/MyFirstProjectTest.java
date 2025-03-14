import org.hyperskill.hstest.dynamic.DynamicTest;
import org.hyperskill.hstest.stage.StageTest;
import org.hyperskill.hstest.testcase.CheckResult;
import org.hyperskill.hstest.testing.TestedProgram;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class MyFirstProjectTest extends StageTest {
    @DynamicTest
    CheckResult testFirstProject() {
        TestedProgram pr = new TestedProgram();

        String output = pr.start().toLowerCase().strip();
        int outputLength = (int) Arrays.stream(output.split("\n")).filter(line -> !line.isEmpty()).count();
        if (output.isEmpty()) {
            return CheckResult.wrong("Your program didn't print any output.");
        } else if (outputLength != 7) {
            return CheckResult.wrong(String.format("Your program should output 7 lines, but %s lines were found.",
                    outputLength));
        }
        if (!output.contains("prices")) {
            return CheckResult.wrong("Your program didn't print the 'Prices:' line");
        }
        Map<String, String> map = new HashMap<>() {{
            put("Bubblegum", "2");
            put("Toffee", "0.2");
            put("Ice cream", "5");
            put("Milk chocolate", "4");
            put("Doughnut", "2.5");
            put("Pancake", "3.2");
        }};
        for (String name : map.keySet()) {
            if (!output.contains(name.toLowerCase())) {
                return CheckResult.wrong(String.format("Your program should print the '%s' as an item", name));
            }
            String expectedPrice = map.get(name);
            if (!output.contains("$" + expectedPrice)) {
                return CheckResult.wrong(String.format("Incorrect or missing dollar sign for '%s' price. " +
                        "Expected output: '%s: $%s'", name, name, expectedPrice));
            }
        }
        return CheckResult.correct();
    }
}
