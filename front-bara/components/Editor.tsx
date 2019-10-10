// const ReactAce = require('react-ace-editor')
import AceEditor from './AceEditor'
import React, { Component } from 'react';

type EditorProps = {
  value?: string;
  onChange: (allValue: string) => void;
}
class CodeEditor extends Component<EditorProps> {
  private ace: any;
  constructor(props: EditorProps) {
    super(props);
    this.onChange = this.onChange.bind(this);
  }
  onChange(newValue: string, e: Event) {
    console.log(newValue, e);

    if(this.ace == null) {
      return
    }
    const editor = this.ace.editor; 
    console.log(editor.getValue())
    this.props.onChange(editor.getValue())
  }
  render() {
    return (
      <AceEditor
        mode="javascript"
        theme="eclipse"
        setReadOnly={false}
        onChange={this.onChange}
        ref={(instance: any) => { this.ace = instance; }}
        setValue={this.props.value}
        option={{fontSize: '14px'}}
      />
    );
  }
}
export default CodeEditor