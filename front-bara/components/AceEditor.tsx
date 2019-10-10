import React, { Component } from "react";
import PropTypes from "prop-types";

let ace;
if (typeof window !== "undefined") {
  ace = require("brace");
}

export type EditorOption = {
    fontSize: string;
}

type Props = {
  editorId?: string;
  onChange?: (newStr: string, e: any) => void;
  option?: EditorOption;
  setReadOnly?: boolean;
  setValue?: string;
  theme?: string;
  mode?: string;
  style?: object;
}

class CodeEditor extends Component<Props> {
  editor: any;
  componentDidMount() {
    if (typeof window !== "undefined") {
      const { onChange, setReadOnly, setValue, theme, mode, option } = this.props;

      require(`brace/mode/${mode}`);
      require(`brace/theme/${theme}`);

      const editor = ace.edit("ace-editor");
      this.editor = editor;
      editor.getSession().setMode(`ace/mode/${mode}`);
      editor.setTheme(`ace/theme/${theme}`);
      editor.on("change", e => onChange(editor.getValue(), e));
      editor.setReadOnly(setReadOnly);
      editor.setValue(setValue);
      editor.setOption(option)
    }
  }

  shouldComponentUpdate() {
    return false;
  }

  render() {
    const { style = {height: '100%', width: '100%'}, editorId = "ace-editor" } = this.props;
    return <div id={editorId} style={style}></div>;
  }
}

export default CodeEditor
