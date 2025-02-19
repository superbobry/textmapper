/**
 * Copyright 2002-2022 Evgeny Gryaznov
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package org.textmapper.templates.ast;

import org.textmapper.templates.api.EvaluationContext;
import org.textmapper.templates.api.IEvaluationStrategy;
import org.textmapper.templates.ast.TemplatesTree.TextSource;

import java.util.List;

public abstract class CompoundNode extends Node {

	protected CompoundNode(TextSource source, int offset, int endoffset) {
		super(source, offset, endoffset);
	}

	protected List<Node> instructions;

	public List<Node> getInstructions() {
		return instructions;
	}

	public void setInstructions(List<Node> instructions) {
		this.instructions = instructions;
	}

	@Override
	protected void emit(StringBuilder sb, EvaluationContext context, IEvaluationStrategy env) {
		if (instructions != null) {
			for (Node n : instructions) {
				n.emit(sb, context, env);
			}
		}
	}

	protected void contentToJavascript(StringBuilder sb) {
		boolean first = true;
		for (Node n : instructions) {
			if (!first) sb.append('+');
			first = false;
			n.toJavascript(sb);
		}
	}
}
