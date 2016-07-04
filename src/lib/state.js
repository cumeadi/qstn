
/**
 * Stateful
 *
 * @param {Object} comp
 * @api public
 */

export default (comp) => {
  const store = {};

  /**
   * Bindings
   */

  const onCreate = m => wrap(comp.onCreate, m);
  const onUpdate = m => wrap(comp.onUpdate, m);
  const render = m => wrap(comp.render, m);

  /**
   * Remove
   *
   * @param {Object} model
   * @api public
   */

  const onRemove = model => {
    delete store[model.path];
    comp.onRemove(model);
  };

  /**
   * Set state
   *
   * @param {Object} obj
   * @api public
   */

  const setState = ({path,dispatch}) => obj => {
    store[path] = {...store[path], ...obj};
    dispatch({type: 'STATE_CHANGE'});
  };

  /**
   * Get state
   *
   * @param {Object} model
   * @api private
   */

  const getState = ({path}) => {
    if(store[path]) return store[path];

    if(comp.initialState) {
      store[path] = comp.initialState();
    } else {
      store[path] = {};
    }

    return store[path];
  };

  /**
   * Wrap
   *
   * @param {Function} fn
   * @param {Object} model
   * @api private
   */

  const wrap = (fn, model) => {
    if('undefined' == typeof fn) return;
    return fn(bind(model));
  };

  /**
   * Bind state
   *
   * @param {Object} model
   * @api private
   */

  const bind = model => ({
    setState: setState(model),
    state: getState(model),
    ...model,
  });

  return {
    ...comp,
    onCreate,
    onUpdate,
    onRemove,
    render,
  };
};

