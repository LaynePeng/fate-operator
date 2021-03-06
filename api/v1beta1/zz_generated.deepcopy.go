// +build !ignore_autogenerated

/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]ImagePullSecrets, len(*in))
		copy(*out, *in)
	}
	out.Istio = in.Istio
	if in.Modules != nil {
		in, out := &in.Modules, &out.Modules
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(Host)
		**out = **in
	}
	if in.Rollsite != nil {
		in, out := &in.Rollsite, &out.Rollsite
		*out = new(Rollsite)
		(*in).DeepCopyInto(*out)
	}
	if in.Lbrollsite != nil {
		in, out := &in.Lbrollsite, &out.Lbrollsite
		*out = new(Lbrollsite)
		(*in).DeepCopyInto(*out)
	}
	if in.Nodemanager != nil {
		in, out := &in.Nodemanager, &out.Nodemanager
		*out = new(Nodemanager)
		(*in).DeepCopyInto(*out)
	}
	if in.Python != nil {
		in, out := &in.Python, &out.Python
		*out = new(Python)
		(*in).DeepCopyInto(*out)
	}
	if in.Mysql != nil {
		in, out := &in.Mysql, &out.Mysql
		*out = new(Mysql)
		(*in).DeepCopyInto(*out)
	}
	if in.Spark != nil {
		in, out := &in.Spark, &out.Spark
		*out = new(Spark)
		(*in).DeepCopyInto(*out)
	}
	if in.Hdfs != nil {
		in, out := &in.Hdfs, &out.Hdfs
		*out = new(Hdfs)
		(*in).DeepCopyInto(*out)
	}
	if in.Nginx != nil {
		in, out := &in.Nginx, &out.Nginx
		*out = new(Nginx)
		(*in).DeepCopyInto(*out)
	}
	if in.Rabbitmq != nil {
		in, out := &in.Rabbitmq, &out.Rabbitmq
		*out = new(Rabbitmq)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Datanode) DeepCopyInto(out *Datanode) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(NodeSelector, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Datanode.
func (in *Datanode) DeepCopy() *Datanode {
	if in == nil {
		return nil
	}
	out := new(Datanode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Exchange) DeepCopyInto(out *Exchange) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Exchange.
func (in *Exchange) DeepCopy() *Exchange {
	if in == nil {
		return nil
	}
	out := new(Exchange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FateCluster) DeepCopyInto(out *FateCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FateCluster.
func (in *FateCluster) DeepCopy() *FateCluster {
	if in == nil {
		return nil
	}
	out := new(FateCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FateCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FateClusterList) DeepCopyInto(out *FateClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FateCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FateClusterList.
func (in *FateClusterList) DeepCopy() *FateClusterList {
	if in == nil {
		return nil
	}
	out := new(FateClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FateClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FateClusterSpec) DeepCopyInto(out *FateClusterSpec) {
	*out = *in
	if in.ClusterSpec != nil {
		in, out := &in.ClusterSpec, &out.ClusterSpec
		*out = new(ClusterSpec)
		(*in).DeepCopyInto(*out)
	}
	out.Kubefate = in.Kubefate
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FateClusterSpec.
func (in *FateClusterSpec) DeepCopy() *FateClusterSpec {
	if in == nil {
		return nil
	}
	out := new(FateClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FateClusterStatus) DeepCopyInto(out *FateClusterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FateClusterStatus.
func (in *FateClusterStatus) DeepCopy() *FateClusterStatus {
	if in == nil {
		return nil
	}
	out := new(FateClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FateJob) DeepCopyInto(out *FateJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FateJob.
func (in *FateJob) DeepCopy() *FateJob {
	if in == nil {
		return nil
	}
	out := new(FateJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FateJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FateJobList) DeepCopyInto(out *FateJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FateJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FateJobList.
func (in *FateJobList) DeepCopy() *FateJobList {
	if in == nil {
		return nil
	}
	out := new(FateJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FateJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FateJobSpec) DeepCopyInto(out *FateJobSpec) {
	*out = *in
	out.FateClusterRef = in.FateClusterRef
	out.JobConf = in.JobConf
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FateJobSpec.
func (in *FateJobSpec) DeepCopy() *FateJobSpec {
	if in == nil {
		return nil
	}
	out := new(FateJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FateJobStatus) DeepCopyInto(out *FateJobStatus) {
	*out = *in
	in.JobStatus.DeepCopyInto(&out.JobStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FateJobStatus.
func (in *FateJobStatus) DeepCopy() *FateJobStatus {
	if in == nil {
		return nil
	}
	out := new(FateJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Fateflow) DeepCopyInto(out *Fateflow) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fateflow.
func (in *Fateflow) DeepCopy() *Fateflow {
	if in == nil {
		return nil
	}
	out := new(Fateflow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hdfs) DeepCopyInto(out *Hdfs) {
	*out = *in
	if in.Namenode != nil {
		in, out := &in.Namenode, &out.Namenode
		*out = new(Namenode)
		(*in).DeepCopyInto(*out)
	}
	if in.Datanode != nil {
		in, out := &in.Datanode, &out.Datanode
		*out = new(Datanode)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hdfs.
func (in *Hdfs) DeepCopy() *Hdfs {
	if in == nil {
		return nil
	}
	out := new(Hdfs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Host) DeepCopyInto(out *Host) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Host.
func (in *Host) DeepCopy() *Host {
	if in == nil {
		return nil
	}
	out := new(Host)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePullSecrets) DeepCopyInto(out *ImagePullSecrets) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePullSecrets.
func (in *ImagePullSecrets) DeepCopy() *ImagePullSecrets {
	if in == nil {
		return nil
	}
	out := new(ImagePullSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Istio) DeepCopyInto(out *Istio) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Istio.
func (in *Istio) DeepCopy() *Istio {
	if in == nil {
		return nil
	}
	out := new(Istio)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobConf) DeepCopyInto(out *JobConf) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobConf.
func (in *JobConf) DeepCopy() *JobConf {
	if in == nil {
		return nil
	}
	out := new(JobConf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kubefate) DeepCopyInto(out *Kubefate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kubefate.
func (in *Kubefate) DeepCopy() *Kubefate {
	if in == nil {
		return nil
	}
	out := new(Kubefate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Kubefate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubefateList) DeepCopyInto(out *KubefateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Kubefate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubefateList.
func (in *KubefateList) DeepCopy() *KubefateList {
	if in == nil {
		return nil
	}
	out := new(KubefateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubefateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubefateNamespacedName) DeepCopyInto(out *KubefateNamespacedName) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubefateNamespacedName.
func (in *KubefateNamespacedName) DeepCopy() *KubefateNamespacedName {
	if in == nil {
		return nil
	}
	out := new(KubefateNamespacedName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubefateSpec) DeepCopyInto(out *KubefateSpec) {
	*out = *in
	in.VolumeSource.DeepCopyInto(&out.VolumeSource)
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubefateSpec.
func (in *KubefateSpec) DeepCopy() *KubefateSpec {
	if in == nil {
		return nil
	}
	out := new(KubefateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubefateStatus) DeepCopyInto(out *KubefateStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubefateStatus.
func (in *KubefateStatus) DeepCopy() *KubefateStatus {
	if in == nil {
		return nil
	}
	out := new(KubefateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Lbrollsite) DeepCopyInto(out *Lbrollsite) {
	*out = *in
	if in.Exchange != nil {
		in, out := &in.Exchange, &out.Exchange
		*out = new(Exchange)
		**out = **in
	}
	if in.PartyList != nil {
		in, out := &in.PartyList, &out.PartyList
		*out = make([]PartyList, len(*in))
		copy(*out, *in)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(NodeSelector, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Lbrollsite.
func (in *Lbrollsite) DeepCopy() *Lbrollsite {
	if in == nil {
		return nil
	}
	out := new(Lbrollsite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *List) DeepCopyInto(out *List) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(NodeSelector, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new List.
func (in *List) DeepCopy() *List {
	if in == nil {
		return nil
	}
	out := new(List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Master) DeepCopyInto(out *Master) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(NodeSelector, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Master.
func (in *Master) DeepCopy() *Master {
	if in == nil {
		return nil
	}
	out := new(Master)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Mysql) DeepCopyInto(out *Mysql) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(NodeSelector, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Mysql.
func (in *Mysql) DeepCopy() *Mysql {
	if in == nil {
		return nil
	}
	out := new(Mysql)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in NGRouteTable) DeepCopyInto(out *NGRouteTable) {
	{
		in := &in
		*out = make(NGRouteTable, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NGRouteTable.
func (in NGRouteTable) DeepCopy() NGRouteTable {
	if in == nil {
		return nil
	}
	out := new(NGRouteTable)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Namenode) DeepCopyInto(out *Namenode) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(NodeSelector, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Namenode.
func (in *Namenode) DeepCopy() *Namenode {
	if in == nil {
		return nil
	}
	out := new(Namenode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Nginx) DeepCopyInto(out *Nginx) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(NodeSelector, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RouteTable != nil {
		in, out := &in.RouteTable, &out.RouteTable
		*out = make(NGRouteTable, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Nginx.
func (in *Nginx) DeepCopy() *Nginx {
	if in == nil {
		return nil
	}
	out := new(Nginx)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in NodeSelector) DeepCopyInto(out *NodeSelector) {
	{
		in := &in
		*out = make(NodeSelector, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSelector.
func (in NodeSelector) DeepCopy() NodeSelector {
	if in == nil {
		return nil
	}
	out := new(NodeSelector)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Nodemanager) DeepCopyInto(out *Nodemanager) {
	*out = *in
	if in.List != nil {
		in, out := &in.List, &out.List
		*out = make([]List, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Nodemanager.
func (in *Nodemanager) DeepCopy() *Nodemanager {
	if in == nil {
		return nil
	}
	out := new(Nodemanager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Party) DeepCopyInto(out *Party) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Party.
func (in *Party) DeepCopy() *Party {
	if in == nil {
		return nil
	}
	out := new(Party)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartyInfo) DeepCopyInto(out *PartyInfo) {
	*out = *in
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = make([]Proxy, len(*in))
		copy(*out, *in)
	}
	if in.Fateflow != nil {
		in, out := &in.Fateflow, &out.Fateflow
		*out = make([]Fateflow, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartyInfo.
func (in *PartyInfo) DeepCopy() *PartyInfo {
	if in == nil {
		return nil
	}
	out := new(PartyInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartyList) DeepCopyInto(out *PartyList) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartyList.
func (in *PartyList) DeepCopy() *PartyList {
	if in == nil {
		return nil
	}
	out := new(PartyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Proxy) DeepCopyInto(out *Proxy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Proxy.
func (in *Proxy) DeepCopy() *Proxy {
	if in == nil {
		return nil
	}
	out := new(Proxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Python) DeepCopyInto(out *Python) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(NodeSelector, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Spark != nil {
		in, out := &in.Spark, &out.Spark
		*out = new(PythonSpark)
		**out = **in
	}
	if in.Hdfs != nil {
		in, out := &in.Hdfs, &out.Hdfs
		*out = new(PythonHdfs)
		**out = **in
	}
	if in.Rabbitmq != nil {
		in, out := &in.Rabbitmq, &out.Rabbitmq
		*out = new(PythonRabbitmq)
		**out = **in
	}
	if in.Nginx != nil {
		in, out := &in.Nginx, &out.Nginx
		*out = new(PythonNginx)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Python.
func (in *Python) DeepCopy() *Python {
	if in == nil {
		return nil
	}
	out := new(Python)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PythonHdfs) DeepCopyInto(out *PythonHdfs) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PythonHdfs.
func (in *PythonHdfs) DeepCopy() *PythonHdfs {
	if in == nil {
		return nil
	}
	out := new(PythonHdfs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PythonNginx) DeepCopyInto(out *PythonNginx) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PythonNginx.
func (in *PythonNginx) DeepCopy() *PythonNginx {
	if in == nil {
		return nil
	}
	out := new(PythonNginx)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PythonRabbitmq) DeepCopyInto(out *PythonRabbitmq) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PythonRabbitmq.
func (in *PythonRabbitmq) DeepCopy() *PythonRabbitmq {
	if in == nil {
		return nil
	}
	out := new(PythonRabbitmq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PythonSpark) DeepCopyInto(out *PythonSpark) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PythonSpark.
func (in *PythonSpark) DeepCopy() *PythonSpark {
	if in == nil {
		return nil
	}
	out := new(PythonSpark)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in RBRouteTable) DeepCopyInto(out *RBRouteTable) {
	{
		in := &in
		*out = make(RBRouteTable, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RBRouteTable.
func (in RBRouteTable) DeepCopy() RBRouteTable {
	if in == nil {
		return nil
	}
	out := new(RBRouteTable)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rabbitmq) DeepCopyInto(out *Rabbitmq) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(NodeSelector, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RouteTable != nil {
		in, out := &in.RouteTable, &out.RouteTable
		*out = make(RBRouteTable, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rabbitmq.
func (in *Rabbitmq) DeepCopy() *Rabbitmq {
	if in == nil {
		return nil
	}
	out := new(Rabbitmq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rollsite) DeepCopyInto(out *Rollsite) {
	*out = *in
	if in.Exchange != nil {
		in, out := &in.Exchange, &out.Exchange
		*out = new(Exchange)
		**out = **in
	}
	if in.PartyList != nil {
		in, out := &in.PartyList, &out.PartyList
		*out = make([]PartyList, len(*in))
		copy(*out, *in)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(NodeSelector, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rollsite.
func (in *Rollsite) DeepCopy() *Rollsite {
	if in == nil {
		return nil
	}
	out := new(Rollsite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Spark) DeepCopyInto(out *Spark) {
	*out = *in
	in.Master.DeepCopyInto(&out.Master)
	in.Worker.DeepCopyInto(&out.Worker)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Spark.
func (in *Spark) DeepCopy() *Spark {
	if in == nil {
		return nil
	}
	out := new(Spark)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Worker) DeepCopyInto(out *Worker) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(NodeSelector, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Worker.
func (in *Worker) DeepCopy() *Worker {
	if in == nil {
		return nil
	}
	out := new(Worker)
	in.DeepCopyInto(out)
	return out
}
