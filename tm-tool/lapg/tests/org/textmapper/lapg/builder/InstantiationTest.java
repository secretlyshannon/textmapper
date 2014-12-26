/**
 * Copyright 2002-2014 Evgeny Gryaznov
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
package org.textmapper.lapg.builder;

import org.junit.Test;
import org.textmapper.lapg.api.TemplateEnvironment;
import org.textmapper.lapg.api.TemplateParameter;
import org.textmapper.lapg.api.TemplateParameter.Type;
import org.textmapper.lapg.api.builder.GrammarBuilder;

import static org.junit.Assert.assertEquals;

public class InstantiationTest {

	@Test
	public void testSuffixes() {
		GrammarBuilder b = GrammarFacade.createBuilder();
		TemplateEnvironment env = GrammarFacade.createBuilder().getRootEnvironment();

		assertEquals("", env.getNonterminalSuffix());

		// Strings
		TemplateParameter s1 = b.addParameter(Type.String, "s1", "abc", null);
		TemplateParameter s2 = b.addParameter(Type.String, "s2", null, null);
		env = env.extend(s1, "abcd");
		assertEquals("_s1-abcd", env.getNonterminalSuffix());

		env = env.extend(s1, "'");
		assertEquals("_s1-Apostrophe", env.getNonterminalSuffix());

		env = env.extend(s1, "q-d123");
		assertEquals("_s1-q-d123", env.getNonterminalSuffix());

		env = env.extend(s1, "");
		assertEquals("_s1-", env.getNonterminalSuffix());

		env = env.extend(s2, "q");
		assertEquals("_s1-_s2-q", env.getNonterminalSuffix());


		// Booleans
		TemplateParameter b1 = b.addParameter(Type.Bool, "b1", null, null);
		env = env.extend(b1, true);
		assertEquals("_b1_s1-_s2-q", env.getNonterminalSuffix());

		env = env.extend(b1, false);
		assertEquals("_nonb1_s1-_s2-q", env.getNonterminalSuffix());

		// Integers
		TemplateParameter i = b.addParameter(Type.Integer, "i", null, null);
		env = env.extend(i, 0);
		assertEquals("_nonb1_i0_s1-_s2-q", env.getNonterminalSuffix());

		env = env.extend(i, 100);
		assertEquals("_nonb1_i100_s1-_s2-q", env.getNonterminalSuffix());

	}

}
