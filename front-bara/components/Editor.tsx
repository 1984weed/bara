// import * as ReactAce from 'react-ace-editor';
const ReactAce = require('react-ace-editor')
import React, { Component } from 'react';

type Props = {}
class CodeEditor extends Component {
  private ace: any;
  constructor(props: Props) {
    super(props);
    this.onChange = this.onChange.bind(this);
  }
  onChange(newValue: string, e: Event) {
    console.log(newValue, e);

    const editor = this.ace.editor; 
    console.log(editor.getValue());
  }
  render() {
    return (
      <ReactAce
        mode="javascript"
        theme="eclipse"
        setReadOnly={false}
        onChange={this.onChange}
        style={{ height: '400px' }}
        ref={(instance: any) => { this.ace = instance; }}
      />
    );
  }
}
export default CodeEditor