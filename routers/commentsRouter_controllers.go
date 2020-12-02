package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["go-youku/controllers:BaseController"] = append(beego.GlobalControllerRouter["go-youku/controllers:BaseController"],
        beego.ControllerComments{
            Method: "Save",
            Router: `/barrage/save`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-youku/controllers:BaseController"] = append(beego.GlobalControllerRouter["go-youku/controllers:BaseController"],
        beego.ControllerComments{
            Method: "BarrageWs",
            Router: `/barrage/ws`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-youku/controllers:BaseController"] = append(beego.GlobalControllerRouter["go-youku/controllers:BaseController"],
        beego.ControllerComments{
            Method: "ChannelRegion",
            Router: `/channel/region`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-youku/controllers:BaseController"] = append(beego.GlobalControllerRouter["go-youku/controllers:BaseController"],
        beego.ControllerComments{
            Method: "ChannelType",
            Router: `/channel/type`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-youku/controllers:CommentController"] = append(beego.GlobalControllerRouter["go-youku/controllers:CommentController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/comment/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-youku/controllers:CommentController"] = append(beego.GlobalControllerRouter["go-youku/controllers:CommentController"],
        beego.ControllerComments{
            Method: "Save",
            Router: `/comment/save`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-youku/controllers:TopController"] = append(beego.GlobalControllerRouter["go-youku/controllers:TopController"],
        beego.ControllerComments{
            Method: "ChannelTop",
            Router: `/channel/top`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-youku/controllers:TopController"] = append(beego.GlobalControllerRouter["go-youku/controllers:TopController"],
        beego.ControllerComments{
            Method: "TypeTop",
            Router: `/type/top`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-youku/controllers:UserController"] = append(beego.GlobalControllerRouter["go-youku/controllers:UserController"],
        beego.ControllerComments{
            Method: "LoginDo",
            Router: `/login/do`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-youku/controllers:UserController"] = append(beego.GlobalControllerRouter["go-youku/controllers:UserController"],
        beego.ControllerComments{
            Method: "SaveRegister",
            Router: `/register/save`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-youku/controllers:VideoController"] = append(beego.GlobalControllerRouter["go-youku/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelAdvert",
            Router: `/channel/advert`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-youku/controllers:VideoController"] = append(beego.GlobalControllerRouter["go-youku/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelHot",
            Router: `/channel/hot`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-youku/controllers:VideoController"] = append(beego.GlobalControllerRouter["go-youku/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelRecommendRegionList",
            Router: `/channel/recommend/region`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-youku/controllers:VideoController"] = append(beego.GlobalControllerRouter["go-youku/controllers:VideoController"],
        beego.ControllerComments{
            Method: "GetChannelRecommendTypeList",
            Router: `/channel/recommend/type`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-youku/controllers:VideoController"] = append(beego.GlobalControllerRouter["go-youku/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelVideo",
            Router: `/channel/video`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-youku/controllers:VideoController"] = append(beego.GlobalControllerRouter["go-youku/controllers:VideoController"],
        beego.ControllerComments{
            Method: "VideoEpisodesList",
            Router: `/video/episodes/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-youku/controllers:VideoController"] = append(beego.GlobalControllerRouter["go-youku/controllers:VideoController"],
        beego.ControllerComments{
            Method: "VideoInfo",
            Router: `/video/info`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
