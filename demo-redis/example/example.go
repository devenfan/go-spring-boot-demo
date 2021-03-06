/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package example

import (
	"github.com/go-spring/go-spring-web/spring-web"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(RedisController))
}

type RedisController struct {
	//RedisTemplate SpringRedis.RedisTemplate `autowire:""`
}

func (c *RedisController) InitWebBean(wc SpringWeb.WebContainer) {
	wc.GET("/get", c.Get)
	wc.POST("/set", c.Set)
}

func (c *RedisController) Get(ctx SpringWeb.WebContext) {
	//key := ctx.FormValue("key")
	//val, err := c.RedisTemplate.Get(ctx, key)
	//if err != nil {
	//	ctx.Error(err)
	//}
	//ctx.JSON(http.StatusOK, val)
}

func (c *RedisController) Set(ctx SpringWeb.WebContext) {
	//key := ctx.FormValue("key")
	//val := ctx.FormValue("val")
	//c.RedisTemplate.Set(ctx, key, val)
	//ctx.NoContent(http.StatusOK)
}
