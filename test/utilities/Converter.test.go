let assert = require('chai').assert;

import { Converter } from '../../src/utilities/Converter';

suite('Converter', ()=> {

    test('longToString', () => {
        assert.equal(null, Converter.longToString(null));
        assert.equal("123", Converter.longToString(123));
    });

    test('stringToLong', () => {
        assert.equal(0, Converter.stringToLong(null, 0));
        assert.equal(0, Converter.stringToLong("ABC", 0));
        assert.equal(123, Converter.stringToLong("123", 0));
    });

    test('doubleToString', () => {
        assert.equal(null, Converter.doubleToString(null));
        assert.equal("123.456", Converter.doubleToString(123.456));
    });

    test('stringToDouble', () => {
        assert.equal(0, Converter.stringToDouble(null, 0));
        assert.equal(0, Converter.stringToDouble("ABC", 0));
        assert.equal(123.456, Converter.stringToDouble("123.456", 0));
    });

    test('booleanToString', () => {
        assert.equal("false", Converter.booleanToString(null));
        assert.equal("true", Converter.booleanToString(true));
    });

    test('stringToBoolean', () => {
        assert.equal(false, Converter.stringToBoolean(null, false));
        assert.equal(true, Converter.stringToBoolean("True", false));
        assert.equal(true, Converter.stringToBoolean("1", false));
        assert.equal(true, Converter.stringToBoolean("T", false));
    });

});
