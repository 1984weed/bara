import AceEditor from "./AceEditor";
import React, { Component } from "react";

type EditorProps = {
    value?: string;
    onChange: (allValue: string) => void;
};
class CodeEditor extends Component<EditorProps> {
    private ace: AceEditor;
    constructor(props: EditorProps) {
        super(props);
        this.onChange = this.onChange.bind(this);
    }

    onChange(newValue: string, e: Event) {
        if (this.ace == null) {
            return;
        }
        const editor = this.ace.editor;
        this.props.onChange(editor.getValue());
    }

    changeDefaultValue(defaultStr: string) {
        this.ace.changeDefaultValue(defaultStr);
    }

    render() {
        return (
            <AceEditor
                mode="javascript"
                theme="eclipse"
                setReadOnly={false}
                onChange={this.onChange}
                ref={(instance: any) => {
                    this.ace = instance;
                }}
                setValue={this.props.value}
                option={{ fontSize: "14px", overwrite: true }}
            />
        );
    }
}
export default CodeEditor;
