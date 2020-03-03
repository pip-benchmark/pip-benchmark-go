let assert = require('chai').assert;

import { Properties } from '../../src/utilities/Properties';

suite('Properties', ()=> {

    test('load', () => {
        let props = new Properties();
        props.loadFromFile('./data/test.properties');

        assert.equal(4, props.lines.length);
        assert.equal('', props['Key1']);
        assert.equal('Value2', props['Key2']);
        assert.equal('"Value 3"', props['Key3']);
    });

    test('save', () => {
        let props = new Properties();
        props.loadFromFile('./data/test.properties');
        props.saveToFile('./data/test_write.properties');
    });

});