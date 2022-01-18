<template>
  <header class="header_all">
    WIKE 压测平台 - 基于go
    <span>项目作者 wike 联系qq 200569525</span>
  </header>
  <div class="container">
    <el-form :model="ruleForm"  ref="ruleFormRef"  >
      <div style="display: flex">
      <el-form-item label=""  prop="method"  style="width:100px;">
        <el-select v-model="ruleForm.method">
          <el-option value="get"></el-option>
          <el-option value="post"></el-option>
          <el-option value="head"></el-option>
          <el-option value="put"></el-option>
          <el-option value="delete"></el-option>
          <el-option value="options"></el-option>
          <el-option value="patch"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="" prop="url" style="flex: 1;margin-left:15px;margin-right:15px">
        <el-input v-model="ruleForm.url"></el-input>
      </el-form-item>
      <el-form-item label="" >
        <el-button type="primary"   @click="submitForm()">立即压测</el-button>
      </el-form-item>
      </div>
      <h2>压测结果</h2>
      <el-divider ></el-divider>
      <pre style="font-size:15px;position: relative;top:-20px" >{{result}}</pre>
      <div>
        <h2>配置请求参数</h2>
        <el-divider></el-divider>
        <el-tabs v-model="activeName" >
          <el-tab-pane label="query参数" name="params">
            <h2>Query  <el-button type="primary"  size="small" style="margin:0 20px"  @click="paramsAdd">添加query参数</el-button></h2>
            <div v-for="(item,index) in form.queryList" style="display: flex;margin:20px 0;padding:0 10px" >
              <el-form-item label="禁用">
                <el-checkbox v-model="item.disable" :key="'disable'+index" label="" size="small"></el-checkbox>
              </el-form-item>
              <el-form-item label="key" :key="'key'+index" style="width:200px;margin-left:15px;margin-right:15px">
                <el-input v-model="item.key"  :disabled="item.disable"  ></el-input>
              </el-form-item>
              <el-form-item label="value" :key="'value'+index"   style="flex: 1;margin-left:15px;margin-right:15px">
                <el-input v-model="item.value" :disabled="item.disable" ></el-input>
              </el-form-item>
              <el-form-item  >
                <el-button type="danger"   @click="paramsDel(index)">删除</el-button>
              </el-form-item>
            </div>
          </el-tab-pane>
          <el-tab-pane label="header参数"  name="header">
            <h2>header  <el-button type="primary"  size="small" style="margin:0 20px"  @click="headerAdd">添加header参数</el-button></h2>
            <div v-for="(item,index) in form.headerList" style="display: flex;margin:20px 0;padding:0 10px" >
              <el-form-item label="禁用">
                <el-checkbox v-model="item.disable" :key="'disable'+index" label="" size="small"></el-checkbox>
              </el-form-item>
              <el-form-item label="key"  :key="'key'+index" style="width:200px;margin-left:15px;margin-right:15px">
                <el-input v-model="item.key"  :disabled="item.disable"  ></el-input>
              </el-form-item>
              <el-form-item label="value" :key="'value'+index"   style="flex: 1;margin-left:15px;margin-right:15px">
                <el-input v-model="item.value" :disabled="item.disable" ></el-input>
              </el-form-item>
              <el-form-item  >
                <el-button type="danger"   @click="headerDel(index)">删除</el-button>
              </el-form-item>
            </div>
          </el-tab-pane>
          <el-tab-pane label="body数据" name="body">

            <el-tabs type="card" v-model="bodyType"  style="padding:10px" >
              <el-tab-pane  name="raw">
                <template #label>
                  <span>
                   <el-radio v-model="bodyType"  label="raw" value="raw"></el-radio>
                  </span>
                </template>
                <el-input v-model="jsonRaw"   type="textarea" :rows="10"  ></el-input>
              </el-tab-pane>
              <el-tab-pane label="" name="form-data">
                <template #label>
                  <span>
                   <el-radio v-model="bodyType"  label="form-data"  value="form-data"></el-radio>
                  </span>
                </template>
                <el-alert style="margin:20px 5px" title="添加和删除字段，选择的文件会重制，请先添加好所有字段，再选择文件"></el-alert>
                <h2>form表单 <el-button type="primary"  size="small" style="margin:0 20px"  @click="bodyAdd">添加body参数</el-button></h2>
                <div v-for="(item,index) in form.bodyList" style="display: flex;margin:20px 0;padding:0 10px" >
                  <el-form-item label="禁用">
                    <el-checkbox v-model="item.disable" :key="'disable'+index" label="" size="small"></el-checkbox>
                  </el-form-item>
                  <el-form-item label="key"  :key="'key'+index" style="width:180px;margin-left:15px;margin-right:15px">
                    <el-input v-model="item.key"  :disabled="item.disable"  ></el-input>
                  </el-form-item>
                  <el-form-item label="数据类型" :key="'type'+index"  style="margin-right:10px"  >
                    <el-select v-model="item.type" style="width:70px;">
                      <el-option value="text"></el-option>
                      <el-option value="file"></el-option>
                    </el-select>
                  </el-form-item>
                  <el-form-item label="value" :key="uuid()"  v-if="item.type=='text'"   style="flex: 1;margin-left:15px;margin-right:15px">
                    <el-input  v-model="item.value" type="text" :disabled="item.disable" ></el-input>
                  </el-form-item>
                  <el-form-item label="value" :key="uuid()"   v-if="item.type=='file'"    style="flex: 1;margin-left:15px;margin-right:15px">
                    <input  type="file"  multiple ref="fileData" :index="index" :disabled="item.disable" />
                  </el-form-item>
                  <el-form-item  >
                    <el-button type="danger"   @click="bodyDel(index)">删除</el-button>
                  </el-form-item>
                </div>
              </el-tab-pane>
            </el-tabs>

          </el-tab-pane>
          <el-tab-pane label="cookie数据" name="cookie">
            <h2>设置cookie 会自动同步到header，只需要把cookie完整贴入就好，下面是例子</h2>
            <el-alert style="margin:20px 5px" title="Cookie_1=value; Path=/; Domain=.www.52wike.com; Expires=Tue, 17 Jan 2023 14:39:03 GMT;"></el-alert>
            <el-input v-model="cookie"   type="textarea" :rows="10"  ></el-input>
          </el-tab-pane>
        </el-tabs>
      </div>
      <div>
        <h2>配置压测参数</h2>
        <el-divider></el-divider>
        <el-form-item label="指定运行的总请求数"  style="flex: 1;margin-left:15px;margin-right:15px">
          <el-input v-model.number="hey.n"></el-input>
        </el-form-item>
        <el-form-item label="客户端并发执行的请求数" style="flex: 1;margin-left:15px;margin-right:15px">
          <el-input v-model.number="hey.c"></el-input>
        </el-form-item>
        <el-form-item label="客户端发送请求的速度限制，以每秒响应数QPS为单位，默认没有限制" style="flex: 1;margin-left:15px;margin-right:15px">
          <el-input v-model.number="hey.q"></el-input>
        </el-form-item>
        <el-form-item label="发送请求的持续时长，超时后程序停止并退出。若指定了持续时间，则忽略总请求数 例如  10s , 5m" style="flex: 1;margin-left:15px;margin-right:15px">
          <el-input v-model="hey.z"></el-input>
        </el-form-item>
        <el-form-item label="每个请求的超时时间（以秒为单位）。默认值为20s，数值0代表永不超时。"  style="flex: 1;margin-left:15px;margin-right:15px">
          <el-input v-model="hey.t"></el-input>
        </el-form-item>
      </div>

    </el-form>
  </div>
</template>

<script>
const api="http://127.0.0.1:1091"
import  parse from 'url-parse'
export default {
  name: 'App',
  components: {
  },
  data:()=>{
    return {
      result:"",
      hey:{
        n:200,
        c:50,
        q:0,
        z:"",
        t:20
      },
      bodyType:"raw",
      ruleForm:{
        method:"get",
        url:"http://www.baidu.com/"
      },
      host:"www.baidu.com",
      jsonRaw:"{}",
      cookie:"lang=zh;",
      form:{
        bodyList:[ {
          key:"",
          value:"",
          disable:false,
          type:"text"
        },{
          key:"",
          value:"",
          disable:false,
          type:"file"
        }
        ],
        queryList:[
          {
            key:"",
            value:"",
            disable:false
          },
          {
            key:"",
            value:"",
            disable:false
          },
          {
            key:"",
            value:"",
            disable:false
          }
        ],
        headerList:[
          {
            key:"Host",
            value:"www.baidu.com",
            disable:false
          },
          {
            key:"User-Agent",
            value:"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36",
            disable:false
          },
          {
            key:"Accept",
            value:"*/*",
            disable:false
          },
          {
            key:"Connection",
            value:"keep-alive",
            disable:false
          },
          {
            key:"set-cookie",
            value:"lang=zh;",
            disable:false
          }
        ]
      },
      activeName:"params"
    }
  },
  methods:{
     uuid(){
       return Math.random()*Math.random()
     },
     paramsDel(index) {
       this.form.queryList.splice(index,1)
     },
     paramsAdd(){
      this.form.queryList.push({
        key:"",
        value:"",
        disable:false
      })
       this.$forceUpdate()
    },
    headerDel(index) {
      this.form.headerList.splice(index,1)
    },
    headerAdd(){
      this.form.headerList.push({
        key:"",
        value:"",
        disable:false
      })
      this.$forceUpdate()
    },
    bodyDel(index){
      this.form.bodyList.splice(index,1)
    },
    bodyAdd(){
      this.form.bodyList.push( {
        key:"",
        value:"",
        disable:false,
        type:"text"
      })
      this.$forceUpdate()
    },
    async submitForm(){
       let load=this.$loading()
       this.result=""
      try {


       function arrToMap(input){
        let obj={}
        for (let i=0;i<input.length;i++){
          obj[input[i].key]=input[i].value
        }
        return obj;
      }
       console.log(this.$refs.fileData)


      if (this.bodyType=="raw"){
        try {
          let bodyData=JSON.parse(this.jsonRaw)
          let  allData={
            "__urlData":this.ruleForm,
            "__data":bodyData,
            "__header":arrToMap(this.form.headerList),
            "__hey":this.hey
          }
          await fetch(api+'/v1/hey?type=json', {
            method: 'POST', // or 'PUT'
            headers: {
            },
            body: JSON.stringify(allData),
          }).then(response => response.text())
              .then(data => {
                this.result=data
                load.close()
              })
        }catch (e) {
          load.close()

          this.$alert("系统异常"+e)
        }
      }else{
        let formData=new FormData()

        this.form.bodyList.filter((item)=>{
          return !item.disable
        }).forEach((item)=>{
          if(item.type=="text"){
            if(item.key&&item.value){
              formData.append(item.key, item.value);
            }
          }
        })
        for (let i=0;i<this.$refs.fileData.length;i++){
          let obj=this.form.bodyList[this.$refs.fileData[i].getAttribute('index')]
          if(!obj.disable&&obj.key){
            if(this.$refs.fileData[i].files.length>1){
              for (let k=0;k<this.$refs.fileData[i].files.length;k++){
                formData.append(obj.key, this.$refs.fileData[i].files[k],this.$refs.fileData[i].files[k].name);
              }
            } if(this.$refs.fileData[i].files.length==1){
                formData.append(obj.key, this.$refs.fileData[i].files[0]);
            }


          }
        }
        formData.append("__urlData", JSON.stringify(this.ruleForm));
        formData.append("__header", JSON.stringify(arrToMap(this.form.headerList)));
        formData.append("__hey",JSON.stringify(this.hey))
        fetch(api+'/v1/hey?type=file', {
          method: 'POST', // or 'PUT'
          headers: {
          },
          body:formData,
        }).then(response => response.text())
            .then(data => {
              this.result=data
              load.close()
            })

      }
      }catch (e) {
         this.$message("系统异常"+e)
        load.close()
      }
    },

  },
  watch:{
    cookie(data){
      for (let j=0;j<this.form.headerList.length;j++){
        if(this.form.headerList[j].key=="set-cookie"){
          this.form.headerList[j]={
            key:'set-cookie',
            value:data,
            disable:false
          }
          break
        }
      }
  },
    ruleForm:{
      handler (item){
        let url = parse(item.url, true);
        this.host=url.host
        for (let j=0;j<this.form.headerList.length;j++){
          if(this.form.headerList[j].key=="Host"){
            this.form.headerList[j]={
              key:'Host',
              value:this.host,
              disable:false
            }
            break
          }
        }
        for (let i in url.query){
          let isNew=true
          for (let j=0;j<this.form.queryList.length;j++){
            if(this.form.queryList[j].key==i){
              isNew=false
              this.form.queryList[j]={
                key:i,
                value:url.query[i],
                disable:false
              }
              break
            }
          }
          if (isNew){
            this.form.queryList.push({
              key:i,
              value:url.query[i],
              disable:false
            })
          }
        }
      },
      deep: true
    },
    form:{
      handler (item){
        let obj={}
        item.queryList.filter((item)=>{
          return !item.disable
        }).forEach(item=>{
          if(item.value&&item.key){
            obj[item.key]=item.value
          }
        })
        let url = parse(this.ruleForm.url, true);

        url.set("query",obj,function (){})
        this.ruleForm.url=url.toString()
      },
      deep: true
      }
  }
}
</script>
