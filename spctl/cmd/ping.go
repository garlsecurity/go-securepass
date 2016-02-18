//    Copyright © 2016 Alessio Treglia <alessio@debian.org>
//
//    This program is free software; you can redistribute it and/or modify
//    it under the terms of the GNU General Public License as published by
//    the Free Software Foundation; either version 2 of the License, or
//    (at your option) any later version.
//
//    This program is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//    GNU General Public License for more details.
//
//    You should have received a copy of the GNU General Public License along
//    with this program; if not, write to the Free Software Foundation, Inc.,
//    51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.

package cmd

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/garlsecurity/go-securepass/securepass"
	"github.com/garlsecurity/go-securepass/spctl/config"
)

func ActionPing(c *cli.Context) {
	s, err := securepass.NewSecurePass(config.Configuration.AppID,
		config.Configuration.AppSecret, config.Configuration.Endpoint)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := s.Ping()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("Ping from IPv%d address %s", resp.IPVersion, resp.IP)
}
