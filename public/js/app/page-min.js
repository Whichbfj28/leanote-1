var em=new editorMode;var lineMove=false;var target=null;function stopResize3Columns(){if(lineMove){ajaxGet("/user/updateColumnWidth",{notebookWidth:UserInfo.NotebookWidth,noteListWidth:UserInfo.NoteListWidth},function(){})}lineMove=false;$(".noteSplit").css("background","none")}function resize3ColumnsEnd(notebookWidth,noteListWidth){if(notebookWidth<150||noteListWidth<100){}var noteWidth=$("body").width()-notebookWidth-noteListWidth;if(noteWidth<400){}$("#leftNotebook").width(notebookWidth);$("#notebookSplitter").css("left",notebookWidth);$("#noteAndEditor").css("left",notebookWidth);$("#noteList").width(noteListWidth);$("#noteSplitter").css("left",noteListWidth);$("#note").css("left",noteListWidth);UserInfo.NotebookWidth=notebookWidth;UserInfo.NoteListWidth=noteListWidth}function resize3Columns(event,isFromeIfr){if(isFromeIfr){event.clientX+=$("body").width()-$("#note").width()}var notebookWidth,noteListWidth;if(lineMove==true){if(target=="notebookSplitter"){notebookWidth=event.clientX;noteListWidth=$("#noteList").width();resize3ColumnsEnd(notebookWidth,noteListWidth)}else{notebookWidth=$("#leftNotebook").width();noteListWidth=event.clientX-notebookWidth;resize3ColumnsEnd(notebookWidth,noteListWidth)}resizeEditor()}}$(function(){$(".noteSplit").bind("mousedown",function(event){event.preventDefault();lineMove=true;$(this).css("background-color","#ccc");target=$(this).attr("id");$("#noteMask").css("z-index",99999)});$("body").bind("mouseup",function(event){stopResize3Columns();$("#noteMask").css("z-index",-1)});$("body").bind("mousemove",function(event){if(lineMove){event.preventDefault();resize3Columns(event)}});$("#moreBtn").click(function(){saveBookmark();var height=$("#mceToolbar").height();if(height<$("#popularToolbar").height()){$("#mceToolbar").height($("#popularToolbar").height());$(this).find("i").removeClass("fa-angle-down").addClass("fa-angle-up")}else{$("#mceToolbar").height(height/2);$(this).find("i").removeClass("fa-angle-up").addClass("fa-angle-down")}resizeEditor();restoreBookmark()});$(window).resize(function(){resizeEditor()});$(".folderHeader").click(function(){var body=$(this).next();var p=$(this).parent();if(!body.is(":hidden")){$(".folderNote").removeClass("opened").addClass("closed");p.removeClass("opened").addClass("closed");$(this).find(".fa-angle-down").removeClass("fa-angle-down").addClass("fa-angle-right")}else{$(".folderNote").removeClass("opened").addClass("closed");p.removeClass("closed").addClass("opened");$(this).find(".fa-angle-right").removeClass("fa-angle-right").addClass("fa-angle-down")}});tinymce.init({setup:function(ed){ed.on("keydown",Note.saveNote);ed.on("keydown",function(e){var num=e.which?e.which:e.keyCode;if(num==9){if(!e.shiftKey){var node=ed.selection.getNode();if(node.nodeName=="PRE"){ed.execCommand("mceInsertRawHTML",false,"	")}else{ed.execCommand("mceInsertRawHTML",false,"&nbsp;&nbsp;&nbsp;&nbsp;")}}else{}e.preventDefault();e.stopPropagation();return false}});ed.on("click",function(e){$("body").trigger("click")});ed.on("click",function(){log(ed.selection.getNode())})},selector:"#editorContent",content_css:["css/bootstrap.css","css/editor/editor.css"].concat(em.getWritingCss()),skin:"custom",language:LEA.locale,plugins:["autolink link leanote_image lists charmap hr","paste","searchreplace leanote_nav leanote_code tabfocus","table directionality textcolor codemirror"],toolbar1:"formatselect | forecolor backcolor | bold italic underline strikethrough | leanote_image | leanote_code | bullist numlist | alignleft aligncenter alignright alignjustify",toolbar2:"outdent indent blockquote | link unlink | table | hr removeformat | subscript superscript |searchreplace | code | pastetext | fontselect fontsizeselect",menubar:false,toolbar_items_size:"small",statusbar:false,url_converter:false,font_formats:"Arial=arial,helvetica,sans-serif;"+"Arial Black=arial black,avant garde;"+"Times New Roman=times new roman,times;"+"Courier New=courier new,courier;"+"Tahoma=tahoma,arial,helvetica,sans-serif;"+"Verdana=verdana,geneva;"+"宋体=SimSun;"+"新宋体=NSimSun;"+"黑体=SimHei;"+"微软雅黑=Microsoft YaHei",block_formats:"Header 1=h1;Header 2=h2;Header 3=h3; Header 4=h4;Pre=pre;Paragraph=p",codemirror:{indentOnInit:true,path:"CodeMirror",config:{lineNumbers:true},jsFiles:[]},paste_data_images:true});window.onbeforeunload=function(e){Note.curChangedSaveIt()};$("body").on("keydown",Note.saveNote)});var random=1;function scrollTo(self,tagName,text){var iframe=$("#editorContent_ifr").contents();var target=iframe.find(tagName+":contains("+text+")");random++;var navs=$('#leanoteNavContent [data-a="'+tagName+"-"+encodeURI(text)+'"]');var len=navs.size();for(var i=0;i<len;++i){if(navs[i]==self){break}}if(target.size()>=i+1){target=target.eq(i);var top=target.offset().top;var nowTop=iframe.scrollTop();var d=200;for(var i=0;i<d;i++){setTimeout(function(top){return function(){iframe.scrollTop(top)}}(nowTop+1*i*(top-nowTop)/d),i)}setTimeout(function(){iframe.scrollTop(top)},d+5);return}}$(function(){$("#leanoteNav h1").on("click",function(e){if(!$("#leanoteNav").hasClass("unfolder")){$("#leanoteNav").addClass("unfolder")}else{$("#leanoteNav").removeClass("unfolder")}});function openSetInfoDialog(whichTab){showDialog("dialogSetInfo",{title:"帐户设置",postShow:function(){$("#myTabs a").eq(whichTab).tab("show");$("#username").val(UserInfo.Username)}})}$("#setInfo").click(function(){if(UserInfo.Email){openSetInfoDialog(0)}else{showDialog("thirdDialogSetInfo",{title:"帐户设置",postShow:function(){$("#thirdMyTabs a").eq(0).tab("show")}})}});$("#setTheme").click(function(){showDialog2("#setThemeDialog",{title:"主题设置",postShow:function(){if(!UserInfo.Theme){UserInfo.Theme="default"}$("#themeForm input[value='"+UserInfo.Theme+"']").attr("checked",true)}})});$("#themeForm").on("click","input",function(e){var val=$(this).val();$("#themeLink").attr("href","/css/theme/"+val+".css");ajaxPost("/user/updateTheme",{theme:val},function(re){if(reIsOk(re)){UserInfo.Theme=val}})});$("#leanoteDialog").on("click","#accountBtn",function(e){e.preventDefault();var email=$("#thirdEmail").val();var pwd=$("#thirdPwd").val();var pwd2=$("#thirdPwd2").val();if(!email){showAlert("#thirdAccountMsg","请输入邮箱","danger","#thirdEmail");return}else{var myreg=/^([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+\.[a-zA-Z]{2,3}$/;if(!myreg.test(email)){showAlert("#thirdAccountMsg","请输入正确的邮箱","danger","#thirdEmail");return}}if(!pwd){showAlert("#thirdAccountMsg","请输入密码","danger","#thirdPwd");return}else{if(pwd.length<6){showAlert("#thirdAccountMsg","密码长度至少6位","danger","#thirdPwd");return}}if(!pwd2){showAlert("#thirdAccountMsg","请重复输入密码","danger","#thirdPwd2");return}else{if(pwd!=pwd2){showAlert("#thirdAccountMsg","两次密码输入不一致","danger","#thirdPwd2");return}}hideAlert("#thirdAccountMsg");post("/user/addAccount",{email:email,pwd:pwd},function(ret){if(ret.Ok){showAlert("#thirdAccountMsg","添加成功!","success");UserInfo.Email=email;$("#curEmail").html(email);hideDialog(1e3)}else{showAlert("#thirdAccountMsg",ret.Msg||"添加失败!","danger")}},this)});$("#leanoteDialog").on("click","#usernameBtn",function(e){e.preventDefault();var username=$("#leanoteDialog #username").val();if(!username){showAlert("#usernameMsg","请输入用户名","danger");return}else if(username.length<4){showAlert("#usernameMsg","用户名长度至少4位","danger");return}else if(/[^0-9a-zzA-Z_\-]/.test(username)){showAlert("#usernameMsg","用户名不能含除数字,字母之外的字符","danger");return}hideAlert("#usernameMsg");post("/user/updateUsername",{username:username},function(ret){if(ret.Ok){UserInfo.UsernameRaw=username;UserInfo.Username=username.toLowerCase();$(".username").html(username);showAlert("#usernameMsg","用户名修改成功!","success")}else{showAlert("#usernameMsg",re.Msg||"该用户名已存在","danger")}},"#usernameBtn")});$("#leanoteDialog").on("click","#emailBtn",function(e){e.preventDefault();var email=isEmailFromInput("#email","#emailMsg");if(!email){return}hideAlert("#emailMsg");post("/user/updateEmailSendActiveEmail",{email:email},function(e){if(e.Ok){var url=getEmailLoginAddress(email);showAlert("#emailMsg","验证邮件已发送, 请及时查阅邮件并验证. <a href='"+url+"' target='_blank'>立即验证</a>","success")}else{showAlert("#emailMsg",e.Msg||"邮件发送失败","danger")}},"#emailBtn")});$("#leanoteDialog").on("click","#pwdBtn",function(e){e.preventDefault();var oldPwd=$("#oldPwd").val();var pwd=$("#pwd").val();var pwd2=$("#pwd2").val();if(!oldPwd){showAlert("#pwdMsg","请输入旧密码","danger","#oldPwd");return}else{if(oldPwd.length<6){showAlert("#pwdMsg","密码长度至少6位","danger","#oldPwd");return}}if(!pwd){showAlert("#pwdMsg","请输入新密码","danger","#pwd");return}else{if(pwd.length<6){showAlert("#pwdMsg","密码长度至少6位","danger","#pwd");return}}if(!pwd2){showAlert("#pwdMsg","请重复输入新密码","danger","#pwd2");return}else{if(pwd!=pwd2){showAlert("#pwdMsg","两次密码输入不一致","danger","#pwd2");return}}hideAlert("#pwdMsg");post("/user/updatePwd",{oldPwd:oldPwd,pwd:pwd},function(e){if(e.Ok){showAlert("#pwdMsg","修改密码成功","success")}else{showAlert("#pwdMsg",e.Msg,"danger")}},"#pwdBtn")});if(!UserInfo.Verified){}$("#wrongEmail").click(function(){openSetInfoDialog(1)});$("#leanoteDialog").on("click",".reSendActiveEmail",function(){showDialog("reSendActiveEmailDialog",{title:"发送验证邮件",postShow:function(){ajaxGet("/user/reSendActiveEmail",{},function(ret){if(typeof ret=="object"&&ret.Ok){$("#leanoteDialog .text").html("发送成功!");$("#leanoteDialog .viewEmailBtn").removeClass("disabled");$("#leanoteDialog .viewEmailBtn").click(function(){hideDialog();var url=getEmailLoginAddress(UserInfo.Email);window.open(url,"_blank")})}else{$("#leanoteDialog .text").html("发送失败")}})}})});$("#leanoteDialog").on("click",".nowToActive",function(){var url=getEmailLoginAddress(UserInfo.Email);window.open(url,"_blank")});$("#notebook, #newMyNote, #myProfile, #topNav, #notesAndSort","#leanoteNavTrigger").bind("selectstart",function(e){e.preventDefault();return false});function updateLeftIsMin(is){ajaxGet("/user/updateLeftIsMin",{leftIsMin:is})}function minLeft(save){$("#leftNotebook").width(30);$("#notebook").hide();$("#noteAndEditor").css("left",30);$("#notebookSplitter").hide();$("#logo").hide();$("#leftSwitcher").hide();$("#leftSwitcher2").show();if(save){updateLeftIsMin(true)}}function maxLeft(save){$("#noteAndEditor").css("left",UserInfo.NotebookWidth);$("#leftNotebook").width(UserInfo.NotebookWidth);$("#notebook").show();$("#notebookSplitter").show();$("#leftSwitcher2").hide();$("#logo").show();$("#leftSwitcher").show();if(save){updateLeftIsMin(false)}}$("#leftSwitcher2").click(function(){maxLeft(true)});$("#leftSwitcher").click(function(){minLeft(true)});function getMaxDropdownHeight(obj){var offset=$(obj).offset();var maxHeight=$(document).height()-offset.top;maxHeight-=70;if(maxHeight<0){maxHeight=0}var preHeight=$(obj).find("ul").height();return preHeight<maxHeight?preHeight:maxHeight}$("#notebookMin div.minContainer").hover(function(){var target=$(this).attr("target");$(this).find("ul").html($(target).html()).show().height(getMaxDropdownHeight(this))},function(){$(this).find("ul").hide()});UserInfo.NotebookWidth=UserInfo.NotebookWidth||$("#notebook").width();UserInfo.NoteListWidth=UserInfo.NoteListWidth||$("#noteList").width();if(LEA.isMobile){UserInfo.NoteListWidth=101}if(UserInfo.LeftIsMin){minLeft(false)}$("#mainMask").html("");$("#mainMask").hide(100);$(".dropdown").on("shown.bs.dropdown",function(){var $ul=$(this).find("ul");$ul.height(getMaxDropdownHeight(this))});$("#tipsBtn").click(function(){showDialog2("#tipsDialog")});$("#yourSuggestions").click(function(){showDialog2("#suggestionsDialog")});$("#suggestionBtn").click(function(e){e.preventDefault();var suggestion=$.trim($("#suggestionTextarea").val());if(!suggestion){$("#suggestionMsg").html("请输入您的建议, 谢谢!").show().addClass("alert-warning").removeClass("alert-success");$("#suggestionTextarea").focus();return}$("#suggestionBtn").html("正在处理...").addClass("disabled");$("#suggestionMsg").html("正在处理...");$.post("/suggestion",{suggestion:suggestion},function(ret){$("#suggestionBtn").html("提交").removeClass("disabled");if(ret.Ok){$("#suggestionMsg").html("谢谢反馈, 我们会第一时间处理, 祝您愉快!").addClass("alert-success").removeClass("alert-warning").show()}else{$("#suggestionMsg").html("出错了").show().addClass("alert-warning").removeClass("alert-success")}})});setTimeout(function(){$("#notebook").slimScroll({height:"100%"});$("#noteItemList").slimScroll({height:"100%"});$("#wmd-input").slimScroll({height:"100%"});$("#wmd-input").css("width","100%");$("#wmd-panel-preview").slimScroll({height:"100%"});$("#wmd-panel-preview").css("width","100%")},10);em.init()});function editorMode(){this.writingHash="#writing";this.normalHash="#normal";this.isWritingMode=location.hash==this.writingHash;this.toggleA=null}editorMode.prototype.toggleAText=function(isWriting){var self=this;setTimeout(function(){toggleA=$("#toggleEditorMode a");if(isWriting){toggleA.attr("href",self.normalHash).text(getMsg("normalMode"))}else{toggleA.attr("href",self.writingHash).text(getMsg("writingMode"))}},0)};editorMode.prototype.isWriting=function(hash){return hash==this.writingHash};editorMode.prototype.init=function(){this.changeMode(this.isWritingMode);var self=this;$("#toggleEditorMode").click(function(){saveBookmark();var $a=$(this).find("a");var isWriting=self.isWriting($a.attr("href"));self.changeMode(isWriting);restoreBookmark()})};editorMode.prototype.changeMode=function(isWritingMode){this.toggleAText(isWritingMode);if(isWritingMode){this.writtingMode()}else{this.normalMode()}$("#moreBtn i").removeClass("fa-angle-up").addClass("fa-angle-down")};editorMode.prototype.resizeEditor=function(){setTimeout(function(){resizeEditor()},10);setTimeout(function(){resizeEditor()},20);setTimeout(function(){resizeEditor()},1e3)};editorMode.prototype.normalMode=function(){var $c=$("#editorContent_ifr").contents();$c.contents().find("#writtingMode").remove();$c.contents().find('link[href$="editor-writting-mode.css"]').remove();$("#noteItemListWrap, #notesAndSort").show();$("#noteList").unbind("mouseenter").unbind("mouseleave");var theme=UserInfo.Theme||"default";theme+=".css";$("#themeLink").attr("href","/css/theme/"+theme);$("#mceToolbar").css("height","30px");this.resizeEditor()};editorMode.prototype.writtingMode=function(){$("#themeLink").attr("href","/css/theme/writting-overwrite.css");setTimeout(function(){var $c=$("#editorContent_ifr").contents();$c.contents().find("head").append('<link type="text/css" rel="stylesheet" href="/css/editor/editor-writting-mode.css" id="writtingMode">')},0);$("#noteItemListWrap, #notesAndSort").fadeOut();$("#noteList").hover(function(){$("#noteItemListWrap, #notesAndSort").fadeIn()},function(){$("#noteItemListWrap, #notesAndSort").fadeOut()});$("#mceToolbar").css("height","40px");this.resizeEditor()};editorMode.prototype.getWritingCss=function(){if(this.isWritingMode){return["css/editor/editor-writting-mode.css"]}return[]};