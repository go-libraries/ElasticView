<template>
  <div class="app-container">

    <el-card class="box-card">
      <el-container>

        <el-main>
          <el-row :gutter="60">
            <el-col :span="12">
              <el-switch v-model="useGorm" active-color="#13ce66" inactive-color="#ff4949" />
              gorm
              <el-switch v-model="useSqlx" active-color="#13ce66" inactive-color="#ff4949" />
              sqlX
              <el-switch v-model="useJson" active-color="#13ce66" inactive-color="#ff4949" />
              json
              <el-switch v-model="useForm" active-color="#13ce66" inactive-color="#ff4949" />
              form
              <el-button type="success" round @click="copy">复制结构体</el-button>
            </el-col>
          </el-row>
          <el-row :gutter="60">
            <el-col :span="12">
              <div>
                <el-form label-position="top" label-width="80px">
                  <el-form-item label="建表SQL">
                    <sql-edit
                      v-model="sqlContent"
                      :simple="true"
                      styles="width: 100%"
                      @getValue="getSql"
                    />

                  </el-form-item>
                </el-form>
              </div>
            </el-col>
            <el-col :span="12">
              <div>
                <el-form label-position="top" label-width="80px">
                  <el-form-item label="go 结构体">
                    <el-input
                      v-model="structContent"
                      spellcheck="false"
                      type="textarea"
                      placeholder="一切顺利的话，这里会有你想要的结果"
                      :rows="30"
                    />
                  </el-form-item>
                </el-form>
              </div>
            </el-col>
          </el-row>
        </el-main>
        <el-footer />
      </el-container>
    </el-card>
  </div>
</template>

<script>

export default {
  name: 'Sql2strut',
  components: {
    SqlEdit: () => import('@/components/SqlEditor/index')
  },
  data() {
    return {
      cache: null,
      sqlContent: ``,
      structContent: '',
      activeIndex: '1',
      typeMap: this.getTypeMap(),
      unsignedTypeMap: this.getUnsignedTypeMap(),
      typeMapStr: '',
      useGorm: true,
      useJson: true,
      useForm: true,
      useSqlx: true,
      dialogFormVisible: false
    }
  },
  watch: {
    sqlContent(val) {
      this.sqlContent = val
      this.translate(this.sqlContent)
    },
    typeMapStr(val) {
      var typeArr = val.split('\n')
      var typeMap = {}
      for (var i = 0, len = typeArr.length; i < len; i++) {
        var itemArr = typeArr[i].split(/\:\s+/)
        if (itemArr[0] != undefined && itemArr[1] != undefined) {
          typeMap[itemArr[0]] = itemArr[1]
        }
      }
      this.typeMap = typeMap
    },
    useSqlx(val) {
      this.useSqlx = val

      this.translate(this.sqlContent)
    },
    useGorm(val) {
      this.useGorm = val

      this.translate(this.sqlContent)
    },
    useJson(val) {
      this.useJson = val

      this.translate(this.sqlContent)
    },
    useForm(val) {
      this.useForm = val
      this.translate(this.sqlContent)
    }
  },
  created() {

  },
  methods: {
    getSql(sql) {
      this.sqlContent = sql
    },
    copy() {
      const v = this.structContent
      const oInput = document.createElement('input')
      oInput.value = v
      document.body.appendChild(oInput)
      oInput.select() // 选择对象;
      document.execCommand('Copy') // 执行浏览器复制命令
      this.$message({
        message: '复制成功',
        type: 'success'
      })
      oInput.remove()
    },
    getUnsignedTypeMap() {
      return {
        'TINYINT': 'uint8',
        'SMALLINT': 'uint16',
        'INT': 'uint',
        'MEDIUMINT': 'uint',
        'BIGINT': 'uint64',
        'FLOAT': 'float64',
        'DOUBLE': 'float64',
        'DECIMAL': 'float64',
        'CHAR': 'string',
        'VARCHAR': 'string',
        'TEXT': 'string',
        'MEDIUMTEXT': 'string',
        'LONGTEXT': 'string',
        'TIME': 'time.Time',
        'DATE': 'time.Time',
        'DATETIME': 'time.Time',
        'TIMESTAMP': 'int64',
        'ENUM': 'string',
        'SET': 'string',
        'BLOB': 'string'
      }
    },
    translate(val) {
      if (!val) {
        this.structContent = ''
        return
      }
      val = val.replace(/ROW_FORMAT\s{0,}=\s{0,}.*?(\s|;)/, '')
      var md
      try {
        md = this.$SqlParser.parse(val)
      } catch (err) {
        this.structContent = '嗯~~，诡异的sql，删掉索引注释和行格式(ROW_FORMAT=***)试试'
        return
      }
      if (md == undefined) {
        this.structContent = '嗯~~，诡异的sql'
      }
      var types = this.typeMap
      var unsignedTypes = this.unsignedTypeMap
      var tableName = this.titleCase(md[0].name) + 'Model'
      var structResult = ''
      for (var i = 0, len = md[0].options.length; i < len; i++) {
        if (md[0].options[i].key == 'COMMENT') {
          structResult += '// ' + md[0].options[i].value + '\n'
        }
      }
      structResult += 'type ' + tableName + ' struct {'
      for (var i = 0, len = md[0].columns.length; i < len; i++) {
        var field = md[0].columns[i]
        if (field.type != 'column') {
          continue
        }
        var fieldName = this.titleCase(field.name)
        console.log(field.data_type)
        if (field.data_type.unsigned) {
          var fieldType = unsignedTypes[field.data_type.type.toLowerCase()]
        } else {
          var fieldType = types[field.data_type.type.toLowerCase()]
        }
        var fieldJsonName = field.name
        console.log(fieldType)
        structResult += '\n\t' + fieldName + ' ' + fieldType + ' '

        var structArr = []
        if (this.useGorm) {
          structArr.push('gorm:"column:' + fieldJsonName + '"')
        }
        if (this.useSqlx) {
          structArr.push('db:"' + fieldJsonName + '"')
        }
        if (this.useJson) {
          structArr.push('json:"' + fieldJsonName + '"')
        }
        if (this.useForm) {
          structArr.push('form:"' + fieldJsonName + '"')
        }
        if (structArr.length > 0) {
          structResult += '`' + structArr.join(' ') + '`'
        }
        if (field.comment.length > 0) {
          structResult += ' //' + field.comment
        }
      }
      structResult += '\n}'
      this.structContent = structResult
    },

    titleCase(str) {
      var array = str.toLowerCase().split('_')
      for (var i = 0; i < array.length; i++) {
        array[i] = array[i][0].toUpperCase() + array[i].substring(1, array[i].length)
      }
      var string = array.join('')

      return string
    },

    getTypeMap() {
      return {
        'tinyint': 'int64',
        'smallint': 'int64',
        'int': 'int64',
        'mediumint': 'int64',
        'bigint': 'int64',
        'float': 'float64',
        'double': 'float64',
        'decimal': 'float64',
        'char': 'string',
        'varchar': 'string',
        'text': 'string',
        'mediumtext': 'string',
        'longtext': 'string',
        'time': 'time.Time',
        'date': 'time.Time',
        'datetime': 'time.Time',
        'timestamp': 'int64',
        'enum': 'string',
        'set': 'string',
        'blob': 'string'
      }
    },

    isJSON(str, pass_object) {
      if (pass_object && this.isObject(str)) return true

      if (!this.isString(str)) return false

      str = str.replace(/\s/g, '').replace(/\n|\r/, '')

      if (/^\{(.*?)\}$/.test(str)) {
        return /"(.*?)":(.*?)/g.test(str)
      }

      if (/^\[(.*?)\]$/.test(str)) {
        return str.replace(/^\[/, '')
          .replace(/\]$/, '')
          .replace(/},{/g, '}\n{')
          .split(/\n/)
          .map(function(s) {
            return this.isJSON(s)
          })
          .reduce(function(prev, curr) {
            return !!curr
          })
      }

      return false
    },

    isString(x) {
      return Object.prototype.toString.call(x) === '[object String]'
    },

    isObject(obj) {
      return Object.prototype.toString.call(obj) === '[object Object]'
    },
    handleSelect(key, keyPath) {
      // console.log(key)
    }

  }
}
</script>

<style scoped>

</style>
