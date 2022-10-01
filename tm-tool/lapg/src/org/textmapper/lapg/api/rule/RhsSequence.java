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
package org.textmapper.lapg.api.rule;

import org.textmapper.lapg.api.ast.AstType;

/**
 * Ordered sequence of elements.
 */
public interface RhsSequence extends RhsRule, RhsPart {

	String getName();

	RhsPart[] getParts();

	RhsMapping getMapping();

	/**
	 * Overrides the type of the mapped field.
	 * Should be a subtype of getMapping().getField().getType();
	 */
	AstType getType();
}
