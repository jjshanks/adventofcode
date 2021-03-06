package net.jjshanks.day3;

import org.junit.Test;

import net.jjshanks.AbstractProblem;
import net.jjshanks.ProblemTestCase;
import net.jjshanks.error.JollyException;

public class Day3Part2Test extends ProblemTestCase {

    @Test
    public void testSimple() throws JollyException {
        AbstractProblem solver = new Day3Part2("day3/simpleInput");
        solver.run();
        assertSolution(336);
    }

}
