const BasePlayer = require("./BasePlayer"); 

cc.Class({
  extends: BasePlayer,

  // LIFE-CYCLE CALLBACKS:
  start() {
    BasePlayer.prototype.start.call(this);
  },

  onLoad() {
    BasePlayer.prototype.onLoad.call(this);
    this.attackedClips = {
      '01': 'attackedTop',
      '0-1': 'attackedBottom',
      '-20': 'attackedLeft',
      '20': 'attackedRight',
      '-21': 'attackedTopLeft',
      '21': 'atackedTopRright',
      '-2-1': 'attackedBottomLeft',
      '2-1': 'attackedBottomRight'
    };
  },

  update(dt) {
    BasePlayer.prototype.update.call(this, dt);
  },

});
