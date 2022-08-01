<template>
  <div v-bind="{ scopedSlots: $scopedSlots }" class="orgchart-container" :style="{height: pageHeight}"
       @wheel="zoom && zoomHandler($event)" @mouseup="pan && panning && panEndHandler($event)">
    <div class="orgchart l2r" :style="{ transform: transformVal, cursor: cursorVal }"
         @mousedown="pan && panStartHandler($event)" @mousemove="pan && panning && panHandler($event)">
      <organization-chart-node :datasource="treeData" :handle-click="handleClick">
        <template v-for="slot in Object.keys($scopedSlots)" :slot="slot" slot-scope="scope">
          <slot :name="slot" v-bind="scope"/>
        </template>
      </organization-chart-node>
    </div>
  </div>
</template>

<script>
  import $ from "jquery";
  import OrganizationChartNode from "./components/node";

  export default {
    name: "OrgChart",
    props: {
      treeData: {
        type: Object,
        required: true,
      },
      pan: {
        type: Boolean,
        required: false,
      },
      zoom: {
        type: Boolean,
        required: false,
      },
      zoomoutLimit: {
        type: Number,
        required: false,
        default: 0.5,
      },
      zoominLimit: {
        type: Number,
        required: false,
        default: 7,
      },
    },
    data() {
      return {
        cursorVal: "default",
        panning: false,
        startX: 0,
        startY: 0,
        transformVal: "",
        pageHeight: "500px"
      };
    },
    components: {
      OrganizationChartNode,
    },
    created() {
      this.setContainerHeight()
    },
    methods: {
      handleClick(nodeData, t) {
        this.$emit("handle-click", nodeData, t);
      },
      setContainerHeight() {
        this.pageHeight = (document.body.clientHeight - 160) + "px"
      },
      panEndHandler() {
        this.panning = false;
        this.cursorVal = "default";
      },
      panHandler(e) {
        let newX = 0;
        let newY = 0;
        if (!e.targetTouches) {
          // pand on desktop
          newX = e.pageX - this.startX;
          newY = e.pageY - this.startY;
        } else if (e.targetTouches.length === 1) {
          // pan on mobile device
          newX = e.targetTouches[0].pageX - this.startX;
          newY = e.targetTouches[0].pageY - this.startY;
        } else if (e.targetTouches.length > 1) {
          return;
        }
        if (this.transformVal === "") {
          if (this.transformVal.indexOf("3d") === -1) {
            this.transformVal = "matrix(1,0,0,1," + newX + "," + newY + ")";
          } else {
            this.transformVal =
              "matrix3d(1,0,0,0,0,1,0,0,0,0,1,0," + newX + ", " + newY + ",0,1)";
          }
        } else {
          let matrix = this.transformVal.split(",");
          if (this.transformVal.indexOf("3d") === -1) {
            matrix[4] = newX;
            matrix[5] = newY + ")";
          } else {
            matrix[12] = newX;
            matrix[13] = newY;
          }
          this.transformVal = matrix.join(",");
        }
      },
      panStartHandler(e) {
        if ($(e.target).closest(".node").length) {
          this.panning = false;
          return;
        } else {
          this.cursorVal = "move";
          this.panning = true;
        }
        let lastX = 0;
        let lastY = 0;
        if (this.transformVal !== "") {
          let matrix = this.transformVal.split(",");
          if (this.transformVal.indexOf("3d") === -1) {
            lastX = parseInt(matrix[4]);
            lastY = parseInt(matrix[5]);
          } else {
            lastX = parseInt(matrix[12]);
            lastY = parseInt(matrix[13]);
          }
        }
        if (!e.targetTouches) {
          // pand on desktop
          this.startX = e.pageX - lastX;
          this.startY = e.pageY - lastY;
        } else if (e.targetTouches.length === 1) {
          // pan on mobile device
          this.startX = e.targetTouches[0].pageX - lastX;
          this.startY = e.targetTouches[0].pageY - lastY;
        } else if (e.targetTouches.length > 1) {
          return;
        }
      },
      setChartScale(newScale) {
        let matrix = "";
        let targetScale = 1;
        if (this.transformVal === "") {
          this.transformVal =
            "matrix(" + newScale + ", 0, 0, " + newScale + ", 0, 0)";
        } else {
          matrix = this.transformVal.split(",");
          if (this.transformVal.indexOf("3d") === -1) {
            targetScale = Math.abs(window.parseFloat(matrix[3]) * newScale);
            if (
              targetScale > this.zoomoutLimit &&
              targetScale < this.zoominLimit
            ) {
              matrix[0] = "matrix(" + targetScale;
              matrix[3] = targetScale;
              this.transformVal = matrix.join(",");
            }
          } else {
            targetScale = Math.abs(window.parseFloat(matrix[5]) * newScale);
            if (
              targetScale > this.zoomoutLimit &&
              targetScale < this.zoominLimit
            ) {
              matrix[0] = "matrix3d(" + targetScale;
              matrix[5] = targetScale;
              this.transformVal = matrix.join(",");
            }
          }
        }
      },
      zoomHandler(e) {
        let newScale = 1 + (e.deltaY > 0 ? -0.2 : 0.2);
        this.setChartScale(newScale);
      },
    },
  };
</script>

<style lang="scss">
  .orgchart-container {
    position: relative;
    display: inline-block;
    width: 100%;
    border-radius: 5px;
    overflow: auto;
    text-align: center;
    background-image: linear-gradient(
            90deg,
            rgba(0, 182, 200, 0.15) 10%,
            rgba(0, 0, 0, 0) 10%
    ),
    linear-gradient(rgba(0, 182, 200, 0.15) 10%, rgba(0, 0, 0, 0) 10%);
    background-size: 10px 10px;
  }

  .orgchart {
    box-sizing: border-box;
    display: inline-block;
    background-size: 10px 10px;
    border: 1px dashed rgba(0, 0, 0, 0);
    padding: 20px;


    table {
      border-spacing: 0;
      border-collapse: separate;
    }

    & > table:first-child {
      margin: 20px auto;
    }

    td {
      text-align: center;
      vertical-align: top;
      padding: 0;
    }

    .lines:nth-child(3) td {
      box-sizing: border-box;
      height: 20px;
    }

    .lines .topLine {
      border-top: 2px solid rgba(24, 144, 255, .8);
    }

    .lines .rightLine {
      border-right: 1px solid rgba(24, 144, 255, .8);
      float: none;
      border-radius: 0;
    }

    .lines .leftLine {
      border-left: 1px solid rgba(24, 144, 255, .8);
      float: none;
      border-radius: 0;
    }

    .lines .downLine {
      background-color: rgba(24, 144, 255, .8);
      margin: 0 auto;
      height: 20px;
      width: 2px;
      float: none;
    }

    /* node styling */
    .node {
      box-sizing: border-box;
      display: inline-block;
      position: relative;
      margin: 0;
      padding: 3px;
      border: 2px dashed transparent;
      text-align: center;
      width: 130px;

      .node-box {
        display: inline-block;

        .operates {
          display: none;
          text-align: center;

          i {
            font-weight: 600;
            cursor: pointer;

            &:not(:first-child) {
              margin-left: 5px;
            }
          }
        }
      }

      &:hover {
        .operates {
          display: block;
        }
      }
    }
  }

  .l2r {
    position: absolute;
    transform: rotate(-90deg) rotateY(-180deg);
    transform-origin: left top;
    margin-left: -50%;

    .node-box {
      transform: rotate(-90deg) translate(-45px, -50px) rotateY(180deg);
      transform-origin: top;
      width: 120px;
    }
  }

  .orgchart.l2r .node {
    width: 65px;
    height: 130px;
    text-align: center;
  }
</style>

<style lang="scss">
  /*.orgchart-container {
    position: relative;
    display: inline-block;
    height: 100%;
    width: 100%; //calc(100% - 24px);
    overflow: auto;
    text-align: center;


    .orgchart {
      box-sizing: border-box;
      display: inline-block;
      -webkit-touch-callout: none;
      -webkit-user-select: none;
      -moz-user-select: none;
      -ms-user-select: none;
      user-select: none;
      background-image: linear-gradient(
              90deg,
              rgba(0, 182, 200, 0.15) 10%,
              rgba(0, 0, 0, 0) 10%
      ),
      linear-gradient(rgba(0, 182, 200, 0.15) 10%, rgba(0, 0, 0, 0) 10%);
      background-size: 10px 10px;
      padding: 0 10px 0;
    }
    .orgchart .verticalNodes > td::before {
      content: "";
      border: 1px solid rgba(24, 144, 255, .8);
    }

    .orgchart table {
      border-spacing: 0;
      border-collapse: separate;
    }

    .orgchart > table:first-child {
      margin: 20px auto;
    }

    .orgchart td {
      text-align: center;
      vertical-align: top;
      padding: 0;
    }

    .orgchart .lines:nth-child(3) td {
      box-sizing: border-box;
      height: 20px;
    }

    .orgchart .lines .topLine {
      border-top: 2px solid rgba(24, 144, 255, .8);
    }

    .orgchart .lines .rightLine {
      border-right: 1px solid rgba(24, 144, 255, .8);
      float: none;
      border-radius: 0;
    }

    .orgchart .lines .leftLine {
      border-left: 1px solid rgba(24, 144, 255, .8);
      float: none;
      border-radius: 0;
    }

    .orgchart .lines .downLine {
      background-color: rgba(24, 144, 255, .8);
      margin: 0 auto;
      height: 20px;
      width: 2px;
      float: none;
    }

    !* node styling *!
    .orgchart .node {
      display: inline-block;
      position: relative;
      margin: 0;
      padding: 2px;
      border: 2px dashed transparent;
      text-align: center;

      .node-box {
        display: inline-block;

        .operates {
          margin-top: 3px;
          display: none;
          text-align: center;

          i {
            font-weight: 600;
            cursor: pointer;

            &:not(:first-child) {
              margin-left: 5px;
            }
          }
        }
      }

      &:hover {
        .operates {
          display: block;
        }
      }
    }
  }*/
</style>