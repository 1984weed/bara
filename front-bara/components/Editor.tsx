// import * as ReactAce from 'react-ace-editor';
const ReactAce = require('react-ace-editor')
import React, { Component } from 'react';

type EditorProps = {
  value?: string;
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
        setValue={this.props.value}
        fontSize={14}
        // setOptions={{
        //   enableBasicAutocompletion: false,
        //   enableLiveAutocompletion: false,
        //   enableSnippets: false,
        //   showLineNumbers: true,
        //   tabSize: 8,
        //   useSoftTabs: true
        // }}
      />
    );
  }
}
export default CodeEditor