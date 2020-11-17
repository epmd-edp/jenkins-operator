## Add Jenkins Slave

Follow the steps below to add a new Jenkins slave:

* Add a new template for Jenkins Slave by navigating to the jenkins-slaves config map under the EDP namespace. Fill in the Key field and add a value:

    ![config-map](../readme-resource/edit_js_configmap.png  "config-map")

_**Note:** To copy an example of template for Jenkins Slave, use this example. The name and label properties should be unique;_

```
<org.csanchez.jenkins.plugins.kubernetes.PodTemplate>
  <inheritFrom></inheritFrom>
  <name>maven</name>
  <namespace></namespace>
  <privileged>false</privileged>
  <alwaysPullImage>false</alwaysPullImage>
  <instanceCap>2147483647</instanceCap>
  <slaveConnectTimeout>100</slaveConnectTimeout>
  <idleMinutes>5</idleMinutes>
  <activeDeadlineSeconds>0</activeDeadlineSeconds>
  <label>maven</label>
  <serviceAccount>jenkins</serviceAccount>
  <nodeSelector></nodeSelector>
  <nodeUsageMode>NORMAL</nodeUsageMode>
  <workspaceVolume class="org.csanchez.jenkins.plugins.kubernetes.volumes.workspace.EmptyDirWorkspaceVolume">
    <memory>false</memory>
  </workspaceVolume>
  <volumes>
    <org.csanchez.jenkins.plugins.kubernetes.volumes.EmptyDirVolume>
      <mountPath>/opt/caches/maven</mountPath>
      <memory>false</memory>
    </org.csanchez.jenkins.plugins.kubernetes.volumes.EmptyDirVolume>
  </volumes>
  <containers>
    <org.csanchez.jenkins.plugins.kubernetes.ContainerTemplate>
      <name>jnlp</name>
      <image>openshift/jenkins-agent-maven-35-centos7:v3.10</image>
      <privileged>false</privileged>
      <alwaysPullImage>false</alwaysPullImage>
      <workingDir>/tmp</workingDir>
      <command></command>
      <args>${computer.jnlpmac} ${computer.name}</args>
      <ttyEnabled>false</ttyEnabled>
      <resourceRequestCpu></resourceRequestCpu>
      <resourceRequestMemory></resourceRequestMemory>
      <resourceLimitCpu></resourceLimitCpu>
      <resourceLimitMemory></resourceLimitMemory>
      <ports/>
    </org.csanchez.jenkins.plugins.kubernetes.ContainerTemplate>
  </containers>
  <envVars/>
  <annotations/>
  <imagePullSecrets/>
  <podRetention class="org.csanchez.jenkins.plugins.kubernetes.pod.retention.Default"/>
</org.csanchez.jenkins.plugins.kubernetes.PodTemplate>
```
    
* Open Jenkins to ensure that everything is added correctly. Click the Manage Jenkins option, navigate to the Configure System menu, and scroll down to the Kubernetes Pod Template with the necessary data: 
![jenkins-slave](../readme-resource/jenkins_k8s_pod_template.png "jenkins-slave")

3. As a result, the newly added Jenkins slave will be available in the Advanced Settings block of the Admin Console tool during the codebase creation:
![advanced-settings](../readme-resource/newly_added_jenkins_slave.png "advanced-settings")