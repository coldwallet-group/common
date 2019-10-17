package log_test

import (
	"testing"
	"github.com/group-coldwallet/common/log"
)


func TestLog(t *testing.T) {
	cfg:= &log.Logcfg{
		Level: log.LvlDebug,
		Env: log.EnvDevelopment,
		TxtType: log.TxtJson,
		TimeKey: "time",
		LevelKey: "level",
		NameKey: "logger",
		CallerKey: "caller",
		MessageKey: "msg",
		StacktraceKey: "stacktrace",
		OutputPaths:  []string{"./service.log","./debug.log"},
		ErrorOutputPaths: []string{"./error.log"},
		//OutputPaths:  nil,
		//ErrorOutputPaths: nil,
	}

	log.InitLog(cfg)

	//单个信息内容
	log.Debug("debug,最灵繁的人也看不见自己的背脊")
	log.Info("info,最困难的事情就是认识自己。")
	log.Warn("warn,有勇气承担命运这才是英雄好汉")
	log.Error("error,与肝胆人共事，无字句处读书。")
	//log.DPanic("dpanic,阅读使人充实，会谈使人敏捷，写作使人精确。")

	//多个信息内容
	log.Debug("debug,最灵繁的人也看不见自己的背脊",1,[]int{1,2,3,4,5},"test")
	log.Info("info,最困难的事情就是认识自己。",1,[]int{1,2,3,4,5},"test")
	log.Warn("warn,有勇气承担命运这才是英雄好汉",1,[]int{1,2,3,4,5},"test")
	log.Error("error,与肝胆人共事，无字句处读书。",1,[]int{1,2,3,4,5},"test")
	//log.DPanic("dpanic,阅读使人充实，会谈使人敏捷，写作使人精确。",1,[]int{1,2,3,4,5},"test")
	//log.Panic("panic,最大的骄傲于最大的自卑都表示心灵的最软弱无力。")
	//log.Fatal("fatal,自知之明是最难得的知识。")
	//log.Panic("panic,最大的骄傲于最大的自卑都表示心灵的最软弱无力。",1,[]int{1,2,3,4,5},"test")
	//log.Fatal("fatal,自知之明是最难得的知识。",1,[]int{1,2,3,4,5},"test")

	//格式化信息
	log.Debugf("debugf,勇气通往天堂，怯懦通往地狱。")
	log.Infof("infof,有时候读书是%s一种巧%s妙地避开思考%s的方法。","test","demo","done")
	log.Warnf("warnf,阅读%s一切好书%s如同和过去%s最杰出的人谈话。","test","demo","done")
	log.Errorf("errorf,越是%s没有本领%s的就越加%s自命不凡。","test","demo","done")
	//log.DPanicf("dpanicf,越是%s无能的人，%s越喜欢挑剔%s别人的错儿。","test","demo","done")
	//log.Panicf("panicf,知人者智%s，自知者明%s。胜人者有力%s，自胜者强。","test","demo","done")
	//log.Fatalf("fatalf,意志坚强%s的人能把%s世界放在手中像泥块%s一样任意揉捏。","test","demo","done")

	//带上多个字段内容
	log.Debugw("debugw,山有木兮木有枝，心悦君兮君不知。",map[string]interface{}{"f1":100000,"d2":"go,go,go","p3":1.12,"l4":[]int{1,2,3,4,5}})
	log.Infow("infow,人生若只如初见，何事秋风悲画扇。",map[string]interface{}{"f1":100000,"d2":"go,go,go","p3":1.12,"l4":[]int{1,2,3,4,5}})
	log.Warnw("warnw,曾经沧海难为水，除却巫山不是云。",map[string]interface{}{"f1":100000,"d2":"go,go,go","p3":1.12,"l4":[]int{1,2,3,4,5}})
	log.Errorw("errorw,人生如逆旅，我亦是行人。",map[string]interface{}{"f1":100000,"d2":"go,go,go","p3":1.12,"l4":[]int{1,2,3,4,5}})
	//log.DPanicw("dpanicw,十年生死两茫茫，不思量，自难忘。",map[string]interface{}{"f1":100000,"d2":"go,go,go","p3":1.12,"l4":[]int{1,2,3,4,5}})
	//log.Panicw("panicw,玲珑骰子安红豆，入骨相思知不知。",map[string]interface{}{"f1":100000,"d2":"go,go,go","p3":1.12,"l4":[]int{1,2,3,4,5}})
	//log.Fatalw("fatalw,只愿君心似我心，定不负相思意。",map[string]interface{}{"f1":100000,"d2":"go,go,go","p3":1.12,"l4":[]int{1,2,3,4,5}})
	log.Sync()
}

