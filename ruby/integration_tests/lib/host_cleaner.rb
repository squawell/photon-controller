# Copyright 2015 VMware, Inc. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"); you may not
# use this file except in compliance with the License. You may obtain a copy of
# the License at http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software distributed
# under the License is distributed on an "AS IS" BASIS, without warranties or
# conditions of any kind, EITHER EXPRESS OR IMPLIED. See the License for the
# specific language governing permissions and limitations under the License.

require "net/ssh"
require_relative "../lib/test_helpers"

module EsxCloud
  class HostCleaner

    class << self

      DATASTORE_DIRS_TO_DELETE = ["disks", "deleted_images", "images", "tmp_images", "vms", "tmp_uploads",
                                  "disk_*", "deleted_image_*", "image_*", "tmp_image_*", "vm_*", "tmp_upload_*"]

      def clean_vms_on_real_host(server, user_name, password)
        dirty_vms = remove_vms server, user_name, password

        Net::SSH.start(server, user_name, {password: password, user_known_hosts_file: "/dev/null"}) do |ssh|
          datastore_dir = "/vmfs/volumes/#{EsxCloud::TestHelpers.get_datastore_name}/"
          rm_cmd = "rm -rf #{datastore_dir}"

          puts "cleaning folders under #{datastore_dir}"
          DATASTORE_DIRS_TO_DELETE.each do |folder_prefix|
            ssh.exec!(rm_cmd + folder_prefix)
          end

          dirty_vms
        end
      end

      def remove_vms(server, user_name, password)
        puts "cleaning vms on host #{server}"
        Net::SSH.start(server, user_name, {password: password, user_known_hosts_file: "/dev/null"}) do |ssh|
          dirty_vms = ssh.exec!("for id in `vim-cmd vmsvc/getallvms | tail -n+2 | awk '{print $1, $2}'`;do echo $id;done")
          ssh.exec!("for id in `vim-cmd vmsvc/getallvms | tail -n+2 | awk '{print $1}'`;do (vim-cmd vmsvc/power.off $id || true) && vim-cmd vmsvc/unregister $id ;done")
          ssh.exec!("tmp=`mktemp` && vim-cmd vmsvc/getallvms 2>$tmp && for id in `awk '{print $4}' $tmp | sed \"s/'//g\"`;do (vim-cmd vmsvc/power.off $id || true) && vim-cmd vmsvc/unregister $id ;done")
          dirty_vms
        end
      end

      def enable_core_dumps(server, user_name, password)
        puts "enabling core dumps on host #{server}"
        Net::SSH.start(server, user_name, {password: password, user_known_hosts_file: "/dev/null"}) do |ssh|
          ssh.exec!("vsish -e set /userworld/global/coreDumpEnabled 1")
        end
      end

      def reboot_host(server, user_name, password)
        puts "rebooting host #{server}"
        Net::SSH.start(server, user_name, {password: password, user_known_hosts_file: "/dev/null"}) do |ssh|
          ssh.exec!("reboot")
        end
      end

      def wait_for_boot(server, user_name, password, max_wait_time_seconds)
        wait_start = Time.now
        puts "waiting for host #{server} to be reachable for #{max_wait_time_seconds} seconds"
        while Time.now - wait_start < max_wait_time_seconds
          begin
            # test if we can ssh into the machine
            Net::SSH.start(server, user_name, {password: password, user_known_hosts_file: "/dev/null", timeout: 5}) do |ssh|
            end
            return
          rescue
          end
        end
        fail "host #{server} did not become available after #{max_wait_time_seconds} seconds"
      end

      def stop_agent(server, user_name, password)
        puts "stopping agent on host #{server}"
        Net::SSH.start(server, user_name, {password: password, user_known_hosts_file: "/dev/null"}) do |ssh|
          ssh.exec!("/etc/init.d/photon-controller-agent stop")
        end
      end

      def start_agent(server, user_name, password)
        puts "starting agent on host #{server}"
        Net::SSH.start(server, user_name, {password: password, user_known_hosts_file: "/dev/null"}) do |ssh|
          ssh.exec!("/etc/init.d/photon-controller-agent start")
        end
      end

      def uninstall_vib(server, user_name, password, vib_name)
        puts "deleting vib #{vib_name} from #{server}"
        Net::SSH.start(server, user_name, {password: password, user_known_hosts_file: "/dev/null"}) do |ssh|
          ssh.exec!("esxcli software vib remove -f -n #{vib_name} | echo #{vib_name}")
        end
      end

      def clean_xenon_state(server, user_name, password)
        puts "cleaning up xenon state and restarting containers on VM #{server}"
        Net::SSH.start(server, user_name, {password: password, user_known_hosts_file: "/dev/null"}) do |ssh|
          ssh.exec!("sudo rm -rf /etc/esxcloud/mustache/cloud-store/cloud-store/sandbox_19000/19000/CloudStoreXenonHost*")
          ssh.exec!("sudo rm -rf /etc/esxcloud/mustache/cloud-store/cloud-store/sandbox_19000/19000/lucene*")
          ssh.exec!("sudo rm -rf /etc/esxcloud/mustache/cloud-store/cloud-store/sandbox_19000/19000/resources/")

          ssh.exec!("sudo rm -rf /etc/esxcloud/mustache/housekeeper/housekeeper/sandbox_16000/16001/HousekeeperXenonServiceHost*")
          ssh.exec!("sudo rm -rf /etc/esxcloud/mustache/housekeeper/housekeeper/sandbox_16000/16001/lucene*")
          ssh.exec!("sudo rm -rf /etc/esxcloud/mustache/housekeeper/housekeeper/sandbox_16000/16001/resources/")

          ssh.exec!("sudo rm -rf /etc/esxcloud/mustache/deployer/deployer/sandbox_18000/18001/DeployerXenonServiceHost*")
          ssh.exec!("sudo rm -rf /etc/esxcloud/mustache/deployer/deployer/sandbox_18000/18001/lucene*")
          ssh.exec!("sudo rm -rf /etc/esxcloud/mustache/deployer/deployer/sandbox_18000/18001/resources/")

          ssh.exec!("sudo rm -rf /etc/esxcloud/mustache/root-scheduler/root-scheduler/sandbox_13010/13010/SchedulerXenonHost*")
          ssh.exec!("sudo rm -rf /etc/esxcloud/mustache/root-scheduler/root-scheduler/sandbox_13010/13010/lucene*")
          ssh.exec!("sudo rm -rf /etc/esxcloud/mustache/root-scheduler/root-scheduler/sandbox_13010/13010/resources/")

          ssh.exec!("docker restart CloudStore Housekeeper Deployer RootScheduler")
        end
      end

      def clean_datastores(server, user_name, password, folders = DATASTORE_DIRS_TO_DELETE)
        puts "cleaning datastores on #{server}"
        Net::SSH.start(server, user_name, {password: password, user_known_hosts_file: "/dev/null"}) do |ssh|
          folders.each do |folder|
            ssh.exec!("for ds in `df | awk '{print $6}' | grep -v Mounted`; do rm -rf $ds/#{folder}; done")
          end
          ssh.exec!("rm -rf /opt/vmware/photon/controller/")
          ssh.exec!("rm -rf /opt/vmware/esxcloud")
        end
      end

      def api_clean(host_ip)
        deployment = EsxCloud::Deployment.find_all.items.first
        host = EsxCloud::Deployment.get_deployment_hosts(deployment.id).items.detect { |h| h.address == host_ip }
        fail "Host with [#{host_ip}] not found." if host.nil?
        enter_suspended_mode host
        EsxCloud::Host.get_host_vms(host.id).items.each { |v| delete_vm v.id }
        enter_maintenance_mode host
        EsxCloud::Host.delete host.id
      end

      def delete_vm(vm_id)
        vm = EsxCloud::Vm.find_vm_by_id vm_id
        stop_vm vm
        vm.disks.each { |d| detach_disk vm, d }
        detach_iso vm
        vm.delete
      end

      def detach_disk(vm, disk)
        if ["persistent-disk", "persistent"].include? disk.kind
          vm.detach_disk disk.id
        end
      end

      def detach_iso(vm)
        begin
          vm.detach_iso
        rescue
        end
      end

      def stop_vm(vm)
        begin
          vm.stop!
        rescue
        end
      end

      def enter_suspended_mode(host)
        begin
          EsxCloud::Host.enter_suspended_mode host.id
        rescue
        end
      end

      def enter_maintenance_mode(host)
        begin
          EsxCloud::Host.enter_maintenance_mode host.id
        rescue
        end
      end

    end

  end
end
