<template>
  <div>
    <div class="json-editor">
      <span style="text-align: center;display:block;" class="font1">
        {{ title }}
        <el-button v-clipboard:copy="value" v-clipboard:success="onCopy" v-clipboard:error="onError" icon="el-icon-document-copy">点击复制</el-button>
      </span>
      <textarea ref="textarea" style="width: 50px" />
    </div>
  </div>
</template>

<script>
import CodeMirror from 'codemirror'
import 'codemirror/addon/lint/lint.css'
import 'codemirror/lib/codemirror.css'
import 'codemirror/theme/rubyblue.css'
import 'codemirror/theme/monokai.css'
import 'codemirror/theme/eclipse.css'
import 'codemirror/theme/idea.css'

import 'codemirror/addon/scroll/annotatescrollbar.js'
import 'codemirror/addon/search/matchesonscrollbar.js'
import 'codemirror/addon/search/match-highlighter.js'
import 'codemirror/addon/search/jump-to-line.js'

import 'codemirror/addon/dialog/dialog.js'
import 'codemirror/addon/dialog/dialog.css'
import 'codemirror/addon/search/searchcursor.js'
import 'codemirror/addon/search/search.js'
import 'codemirror/addon/fold/foldgutter.css'
import 'codemirror/addon/fold/foldcode'
import 'codemirror/addon/fold/foldgutter'
import 'codemirror/addon/fold/brace-fold'
import 'codemirror/addon/fold/comment-fold'

import 'codemirror/mode/javascript/javascript'
import 'codemirror/addon/lint/lint'
import 'codemirror/addon/lint/json-lint'
import 'codemirror/addon/scroll/simplescrollbars.css'
import 'codemirror/addon/scroll/simplescrollbars'

require('script-loader!jsonlint')

export default {
  name: 'JsonEditor',
  /* eslint-disable vue/require-prop-types */
  props: {
    value: {
      type: String,
      default: '{}'
    },
    read: {
      type: Boolean,
      default: false
    },
    title: {
      type: String
    }
  },
  data() {
    return {
      jsonEditor: false
    }
  },
  watch: {
    value(value) {
      const editorValue = this.jsonEditor.getValue()
      if (value !== editorValue) {
        // var tempJson = JSON.parse(this.value)
        // this.jsonEditor.setValue(JSON.stringify(tempJson, null, '\t'), null, 2)
        this.jsonEditor.setValue(this.value)
      }
    }
  },
  mounted() {
    console.log(this.title)
    this.jsonEditor = CodeMirror.fromTextArea(this.$refs.textarea, {
      gutters: ['CodeMirror-linenumbers', 'CodeMirror-foldgutter', 'CodeMirror-lint-markers'],
      lineNumbers: true,
      mode: 'application/json',
      theme: 'idea',
      lineWrapping: true,
      lint: true,
      scrollbarStyle: 'simple',
      smartIndent: true,
      foldGutter: true,
      readOnly: this.read,
      gutters: ['CodeMirror-lint-markers'],
      extraKeys: {
        'F7': function autoFormat(editor) {
          const value = editor.getValue()
          if (value === '') return
          var tempJson = {}
          try {
            tempJson = JSON.parse(value)
          } catch (error) {
            return
          }
          editor.setValue(JSON.stringify(tempJson, null, '\t'))
        }// 代码格式化
      }
    })

    this.jsonEditor.on('change', cm => {
      this.$emit('changed', cm.getValue())
      this.$emit('input', cm.getValue())
    })

    if (this.value == '') {
      this.value = '{}'
    }
    var tempJson = {}
    try {
      tempJson = JSON.parse(this.value)
    } catch (error) {

    } finally {
      this.jsonEditor.setValue(JSON.stringify(tempJson, null, '\t'))
    }
  },
  destroyed() {

  },
  methods: {
    onCopy(e) { 　　 // 复制成功
      this.$message({
        message: '复制成功！',
        type: 'success'
      })
    },
    onError(e) {　　 // 复制失败
      this.$message({
        message: '复制失败！',
        type: 'error'
      })
    },
    getValue() {
      console.log(this.jsonEditor.getValue(), 'this.jsonEditor.getValue()')
      return this.jsonEditor.getValue()
    },
    isJson(s) {
      if (s === '') return true
      try {
        JSON.parse(s)
        return true
      } catch (error) {
        return false
      }
    }
  }
}
</script>

<style scoped>
  .font1{
    font-family: "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif;
    color: green;
  }
  .json-editor {
    height: 800px;
    width: 49%;
    float: left;
    position: relative;
  }

  .json-editor >>> .CodeMirror {
    min-height: 800px;
  }

  .json-editor >>> .CodeMirror-scroll {
    min-height: 300px;
  }

  .json-editor >>> .cm-s-rubyblue span.cm-string {
    color: #F08047;
  }
</style>
