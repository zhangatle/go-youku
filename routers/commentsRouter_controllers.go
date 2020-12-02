package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

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

}
