let assert = require('chai').assert;

import { PropertyFileLine } from '../../src/utilities/PropertyFileLine';

suite('PropertyFileLine', ()=> {

    test('compose', () => {
        let line = new PropertyFileLine("Key", "Value", "Comment");
        assert.equal("Key", line.key);
        assert.equal("Value", line.value);
        assert.equal("Comment", line.comment);
        assert.equal("Key=Value ;Comment", line.line);
    });

    test('parse key', () => {
        let line = new PropertyFileLine("Key");
        assert.equal("Key", line.key);
        assert.equal("", line.value);
        assert.equal(null, line.comment);
    });

    test('parse key/value', () => {
        let line = new PropertyFileLine("Key=Value");
        assert.equal("Key", line.key);
        assert.equal("Value", line.value);
        assert.equal(null, line.comment);

        line = new PropertyFileLine("Key='Value'");
        assert.equal("Key", line.key);
        assert.equal("Value", line.value);
        assert.equal(null, line.comment);

        line = new PropertyFileLine('Key="Value"');
        assert.equal("Key", line.key);
        assert.equal("Value", line.value);
        assert.equal(null, line.comment);
    });

    test('parse full line', () => {
        let line = new PropertyFileLine("Key=Value;Comment");
        assert.equal("Key", line.key);
        assert.equal("Value", line.value);
        assert.equal("Comment", line.comment);
    });

});
