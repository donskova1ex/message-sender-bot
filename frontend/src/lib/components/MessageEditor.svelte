<script lang="ts">
    import { onMount, onDestroy } from "svelte";
    import type { Editor } from "@tiptap/core";
    import { Editor as TiptapEditor } from "@tiptap/core";
    import StarterKit from "@tiptap/starter-kit";

    export let content = '';

    let editor: Editor | null = null
    let editorElement: HTMLDivElement | null = null;

    onMount(() => {
        if (!editorElement) return;

        editor = new TiptapEditor({
            element: editorElement,
            extensions: [StarterKit],
            content,
            autofocus: false,
        });
    });
    onDestroy(() => {
        if (editor) {
            editor.destroy();
            editor = null;
        }
        });

  export function getHTML() {
    return editor?.getHTML() || '';
  }

  export function getText() {
    return editor?.getText({ blockSeparator: '\n' }) || '';
  }
</script>