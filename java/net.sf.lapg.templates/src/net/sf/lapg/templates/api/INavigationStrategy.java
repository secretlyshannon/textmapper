package net.sf.lapg.templates.api;

/**
 * Navigation strategy specifies how to take a property or call a method of an object. It can be registered for classes
 * or concrete instances.
 */
public interface INavigationStrategy {

	/**
	 * Returns value of <propertyName> property of <obj>.
	 */
	public Object getProperty(Object obj, String propertyName) throws EvaluationException;

	/**
	 * Returns a result of <obj>.methodName(args) call.
	 */
	public Object callMethod(Object obj, String methodName, Object[] args) throws EvaluationException;

	/**
	 * Returns indexed value.
	 */
	public Object getByIndex(Object obj, Object index) throws EvaluationException;

	/**
	 * Factory returns strategy for objects.
	 */
	public static interface Factory {

		/**
		 * Connects factory to the evaluation strategy.
		 */
		public void setEvaluationStrategy(IEvaluationStrategy strategy);

		/**
		 * Returns navigation strategy for object instance.
		 */
		public INavigationStrategy getStrategy(Object o);
	}
}
