function Link(obj){
    this.obj = obj;
    this.link = this.obj.currentTarget.href;
}
Link.prototype.type = "链接";
Link.prototype.tip = function(){
  this.obj.preventDefault();
  if(confirm("删除是不可恢复的，你确认要删除吗？")){
    location.href = this.link;
  }else{
    console.log('取消了');
  }
};
