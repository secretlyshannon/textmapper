package net.sf.lapg.templates.ast;

import net.sf.lapg.templates.ExecutionEnvironment;

public class ThisNode extends ExpressionNode {

	public Object resolve(Object context, ExecutionEnvironment env) {
		return context;
	}
}
