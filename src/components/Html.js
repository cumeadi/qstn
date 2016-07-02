
import {element} from 'deku';

/**
 * Render
 *
 * @param {Object} obj
 * @return {element}
 * @api public
 */

function render({children}) {
  return (
    <html>
    <meta charset="utf-8" />
    <title>?!</title>
    <script src="/bundle.js" defer></script>
    <body>
    {children}
    </body>
    </html>
  );
}

/**
 * Expose
 */

export default {
  render,
};

