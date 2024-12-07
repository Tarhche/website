"use client";
import {useImperativeHandle, forwardRef} from "react";
import {RichTextEditor, Link} from "@mantine/tiptap";
import Highlight from "@tiptap/extension-highlight";
import StarterKit from "@tiptap/starter-kit";
import Underline from "@tiptap/extension-underline";
import TextAlign from "@tiptap/extension-text-align";
import Superscript from "@tiptap/extension-superscript";
import SubScript from "@tiptap/extension-subscript";
import CodeBlockLowlight from "@tiptap/extension-code-block-lowlight";
import {useEditor} from "@tiptap/react";
import {createLowlight} from "lowlight";
import go from "highlight.js/lib/languages/go";
import "@mantine/tiptap/styles.css";
import "./tiptap-editor.css";

const lowlight = createLowlight();
lowlight.register("go", go);

type Props = {
  defaultValue?: string;
};

export type EditorRef = {
  getHTML: (() => string) | undefined;
};

export const TipTapEditor = forwardRef<EditorRef, Props>((props, ref) => {
  const {defaultValue} = props;

  const editor = useEditor({
    extensions: [
      Underline,
      Link,
      Superscript,
      SubScript,
      Highlight,
      StarterKit.configure({codeBlock: false}),
      CodeBlockLowlight.configure({
        lowlight,
      }),
      TextAlign.configure({types: ["heading", "paragraph"]}),
    ],
    content: defaultValue,
    immediatelyRender: false,
  });

  useImperativeHandle(ref, () => {
    return {
      getHTML: editor?.getHTML.bind(editor),
    };
  }, [editor]);

  return (
    <RichTextEditor editor={editor}>
      <RichTextEditor.Toolbar sticky stickyOffset={60}>
        <RichTextEditor.ControlsGroup>
          <RichTextEditor.Bold />
          <RichTextEditor.Italic />
          <RichTextEditor.Underline />
          <RichTextEditor.Strikethrough />
          <RichTextEditor.ClearFormatting />
          <RichTextEditor.Highlight />
          <RichTextEditor.CodeBlock />
        </RichTextEditor.ControlsGroup>
        <RichTextEditor.ControlsGroup>
          <RichTextEditor.H1 />
          <RichTextEditor.H2 />
          <RichTextEditor.H3 />
          <RichTextEditor.H4 />
        </RichTextEditor.ControlsGroup>
        <RichTextEditor.ControlsGroup>
          <RichTextEditor.Blockquote />
          <RichTextEditor.Hr />
          <RichTextEditor.BulletList />
          <RichTextEditor.OrderedList />
          <RichTextEditor.Subscript />
          <RichTextEditor.Superscript />
        </RichTextEditor.ControlsGroup>
        <RichTextEditor.ControlsGroup>
          <RichTextEditor.Link />
          <RichTextEditor.Unlink />
        </RichTextEditor.ControlsGroup>
        <RichTextEditor.ControlsGroup>
          <RichTextEditor.AlignLeft />
          <RichTextEditor.AlignCenter />
          <RichTextEditor.AlignJustify />
          <RichTextEditor.AlignRight />
        </RichTextEditor.ControlsGroup>
        <RichTextEditor.ControlsGroup>
          <RichTextEditor.Undo />
          <RichTextEditor.Redo />
        </RichTextEditor.ControlsGroup>
      </RichTextEditor.Toolbar>
      <RichTextEditor.Content />
    </RichTextEditor>
  );
});

TipTapEditor.displayName = "TipTapEditor";