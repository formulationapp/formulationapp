<!-- Adapted from https://tiptap.dev/installation/vue3 -->
<template>
  <editor-content :editor="editor" spellcheck="false"
                  @keyup.enter.capture.prevent="() => {}"
                  @keydown.enter.capture.prevent="() => {}"/>
</template>

<script setup lang="ts">
import {computed, watch} from 'vue'
import Document from '@tiptap/extension-document'
import Paragraph from '@tiptap/extension-paragraph'
import Text from '@tiptap/extension-text'
import Bold from '@tiptap/extension-bold'
import Italic from '@tiptap/extension-italic'
import History from '@tiptap/extension-history'
import Placeholder from '@tiptap/extension-placeholder'
import Link from '@tiptap/extension-link'
import {EditorContent, useEditor} from '@tiptap/vue-3'
import {htmlToMarkdown, markdownToHtml} from "@/lib/lotion";

const props = defineProps({
  modelValue: {
    type: String,
    default: '',
  },
  placeholder: {
    type: String,
    default: 'Type something...'
  }
})

const emit = defineEmits(['update:modelValue'])

const value = computed({
  get() {
    const mdValue = props.modelValue
    if (mdValue) {
      return markdownToHtml(mdValue)
    } else {
      return null;
    }
  },
  set(newValue) {
    emit('update:modelValue', newValue)
  },
})

const editor = useEditor({
  extensions: [
    Document,
    Paragraph,
    Text,
    Bold,
    Italic,
    History,
    Link,
    Placeholder.configure({
      placeholder: props.placeholder
    })
  ],
  editorProps: {
    // Removing default behavior for drop event
    handleDrop: () => true,
  },
  content: value.value,
  onUpdate: () => {
    value.value = htmlToMarkdown(editor.value?.getHTML() || '')
  },
})

watch(() => props.modelValue, value => {
  const isSame = htmlToMarkdown(editor.value?.getHTML() || '') === value
  if (isSame) return
  editor.value?.commands.setContent(markdownToHtml(value), false)
})
</script>

<style>
.lotion .ProseMirror p.is-editor-empty:first-child::before {
  content: attr(data-placeholder);
  float: left;
  color: #adb5bd;
  pointer-events: none;
  height: 0;
}
</style>